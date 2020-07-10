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
	"github.com/INFURA/eth2-comply/pkg/testcases"
)

func main() {
	// Setup and parse CLI arguments.
	testsRoot := flag.String("testsRoot", "", "Path to a directory tree with test cases")
	testsRemote := flag.String("testsRemote", "https://github.com/INFURA/eth2-comply/releases/download/v0.1.0/tests-v0.1.0.zip", "URL of a ZIP file containing a directory tree with test cases")
	outDir := flag.String("outDir", "/tmp", "A directory where zip files will be downloaded and unzipped.")
	target := flag.String("target", "NO TARGET PROVIDED", "A URL to run tests against, for example http://localhost:5051")
	timeout := flag.String("timeout", "10m", "The time to wait for a case execution to complete. For example, 3600s, 60m, 1h")
	flag.Parse()

	// Setup OAPI client.
	targetUrl, err := url.Parse(*target)
	if err != nil {
		panic(err)
	}
	oapiClient := oapi.NewClient(*targetUrl)

	// Get test cases.
	opts := &testcases.TestsCasesOpts{
		Target:      *target,
		TestsRoot:   *testsRoot,
		TestsRemote: *testsRemote,
		OutDir:      *outDir,
		OapiClient:  oapiClient,
	}
	testCases, err := testcases.All(opts)
	if err != nil {
		panic(err)
	}

	// Create test context with timeout.
	timeoutDur, err := time.ParseDuration(*timeout)
	if err != nil {
		panic(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeoutDur)

	// Execute test cases.
	for _, testCase := range testCases {
		go testCase.Exec(ctx)
	}

	// Concurrent helper function cancels contexts (tests) that timeout.
	go func(ctx context.Context, cancelFunc func()) {
		deadline, ok := ctx.Deadline()
		if !ok {
			panic("Could not get deadline for background context. Is your timeout parameter specified properly?")
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

	// Print test results as they come in.
	for _, testCase := range testCases {
		if testCase.Config.AwaitSlot > 0 {
			fmt.Printf("Next test is awaiting slot %d\n", testCase.Config.AwaitSlot)
		}
		<-testCase.Done
		testCase.PrintResults()
	}

	// If any test was unsuccessful, exit with code 1.
	for _, testCase := range testCases {
		if !testCase.Result.Success {
			os.Exit(1)
		}
	}
}
