package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/INFURA/eth2-comply/pkg/oapi"
	"github.com/INFURA/eth2-comply/pkg/target"
	"github.com/INFURA/eth2-comply/pkg/testcases"
)

func main() {
	// Setup and parse CLI arguments.
	testsRoot := flag.String("testsRoot", "", "Path to a directory tree with test cases")
	testsRemote := flag.String("testsRemote", "https://github.com/INFURA/eth2-comply/releases/download/v0.3.1/tests-v0.3.1.zip", "URL of a ZIP file containing a directory tree with test cases")
	outDir := flag.String("outDir", "/tmp", "A directory where zip files will be downloaded and unzipped.")
	targetLoc := flag.String("target", "NO TARGET PROVIDED", "A URL to run tests against, for example http://localhost:5051")
	timeout := flag.String("timeout", "10s", "The time to wait for a case execution to complete. For example, 3600s, 60m, 1h")
	subset := flag.String("subset", "/", "The subset of paths to run tests for. For example, set this to \"/v1/node\" to only run tests for routes in that path. Defaults to \"/\" (all paths).")
	failSilent := flag.Bool("failSilent", false, "When true, return a 0 code even when tests fail. Defaults to false.")
	flag.Parse()

	// Create test context with timeout.
	timeoutDur, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeoutDur)

	// Setup OAPI client.
	targetUrl, err := url.Parse(*targetLoc)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	ctx = oapi.WithClient(ctx, *targetUrl)

	// Wait for target to become healthy
	if err := target.IsHealthy(ctx); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// Get test cases.
	opts := &testcases.TestsCasesOpts{
		Target:      *targetLoc,
		TestsRoot:   *testsRoot,
		TestsRemote: *testsRemote,
		OutDir:      *outDir,
	}
	testCases, err := testcases.All(opts)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// Execute test cases.
	for _, testCase := range testCases {
		go testCase.Exec(ctx, *subset)
	}

	// Concurrent helper function cancels contexts (tests) that timeout.
	go func(ctx context.Context, cancelFunc func()) {
		deadline, ok := ctx.Deadline()
		if !ok {
			fmt.Printf("Could not get deadline for background context. Is timeout parameter specified properly?\n")
			os.Exit(1)
		}
		for {
			if time.Now().After(deadline) {
				cancelFunc()
				return
			}
			time.Sleep(time.Second)
		}
	}(ctx, cancelFunc)

	// Sort testCases by awaitSlot. This enables us to always print the latest
	// available test result.
	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].Config.AwaitSlot < testCases[j].Config.AwaitSlot
	})

	// Print test results as they come in, and record whether any test failed.
	hasFailures := false
	for _, testCase := range testCases {
		<-testCase.Done

		fmt.Printf("%s\n", testCase.ResultsPretty())

		if !testCase.Result.Success {
			hasFailures = true
		}
	}

	// If any test was unsuccessful, exit with code 1.
	if hasFailures && !*failSilent {
		os.Exit(1)
	}
}
