// package testcases defines the Case data structure, and a public function for
// instantiating slices of Cases from JSON test specifications in directory
// trees on the filesystem.
package testcases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/INFURA/eth2-comply/pkg/eth2spec"
	"github.com/INFURA/eth2-comply/pkg/oapi"
)

// Case is an executable test case. The Config property can be accessed to get
// information about the case scenario.
type Case struct {
	Config     CaseConfig
	OapiClient *eth2spec.APIClient
	Result     Result
	Skipped    bool
	Done       chan struct{}
}

// CaseConfig describes a test scenario.
type CaseConfig struct {
	Method             string
	Route              string
	AwaitSlot          int
	QueryParams        map[string]string
	ReqBody            interface{}
	ExpectedRespStatus int
	ExpectedRespBody   interface{}
}

// Result describes the result of a test. Error is nil is success is true.
type Result struct {
	Success bool
	Error   error
}

// NewCase instantiates and returns a Case struct.
func NewCase(config CaseConfig, oapiClient *eth2spec.APIClient) *Case {
	c := &Case{}

	c.Config = config
	c.OapiClient = oapiClient
	c.Result = Result{}
	c.Done = make(chan struct{})

	return c
}

type OapiError struct {
	Err            error
	ServerResponse []byte
}

func (e OapiError) Error() string {
	return fmt.Sprintf("  OpenAPI client error! A 404 error means the target does not implement the route. A json unmarshaling error means the target implements the route response incorrectly.\n\n  Error:\n    %s\n  Received server message:\n    %s", e.Err.Error(), string(e.ServerResponse))
}

type ExpectationsError struct {
	Err error
}

func (e ExpectationsError) Error() string {
	return fmt.Sprintf("  Response did not satisfy expectations!\n\n  Error:\n%s", e.Err.Error())
}

// Exec executes a test Case and populates the Case's Result struct. The Result
// is unsuccessful and an error is stored in any of three cases:
//
// 1. The response was invalid for the request according to the OAPI schema.
// 2. The response contents did not satisfy the Case's expectations (if any).
// 3. Other reasons pertaining to the case specification or the environment.
//    For example, the case contains an invalid route or cannot be unmarshaled
//    (is ill-formed), or a network condition prevented contacting the target.
//
// Otherwise, the Result is marked as a Success and the Error is left nil.
func (c *Case) Exec(ctx context.Context, pathsRoot string) {
	defer close(c.Done)

	// If a test should be excluded because it is not beneath the paths root,
	// skip it here.
	if !strings.HasPrefix(c.Config.Route, pathsRoot) {
		c.Skipped = true
		return
	}

	// Otherwise, just wait for the node to be healthy.
	err := c.awaitTargetIsHealthy(ctx)
	if err != nil {
		c.setFailure(err)
		return
	}

	// If a test specifies an await slot, wait for the node to sync that slot.
	if c.Config.AwaitSlot > 0 {
		err = c.awaitTargetHasSlot(ctx)
		if err != nil {
			c.setFailure(err)
			return
		}
	}

	result, err := c.execOperation(ctx)
	if err != nil {
		// If the response is invalid in the OAPI schema, set that error here.
		if oapiErr, ok := err.(eth2spec.GenericOpenAPIError); ok {
			if len(oapiErr.Body()) > 0 {
				c.setFailure(OapiError{Err: oapiErr, ServerResponse: oapiErr.Body()})
				return
			}
		}

		// If an environmental error like a network failure occurred, set that
		// failure here.
		c.setFailure(err)
		return
	}

	err = c.assertExpectations(result)
	if err != nil {
		c.setFailure(ExpectationsError{Err: err})
		return
	}

	c.Result.Success = true
}

// setFailure marks a test case as having failed and records a corresponding
// error.
func (c *Case) setFailure(err error) {
	c.Result.Success = false
	c.Result.Error = err
}

// PrintResults pretty-prints a test case and its result to stdout.
func (c Case) PrintResults() {
	route, err := url.Parse(c.Config.Route)
	if err != nil {
		// Panic on user-error
		panicMsg := fmt.Sprintf("Could not parse url %s when printing test results", c.Config.Route)
		panic(panicMsg)
	}

	routeString := fmt.Sprintf("%s %s", c.Config.Method, route.RequestURI())
	if len(c.Config.QueryParams) > 0 {
		routeString = fmt.Sprintf("%s%s", routeString, "?")
		for i, x := range c.Config.QueryParams {
			routeString = fmt.Sprintf("%s%s=%s,", routeString, i, x)
		}
		routeString = routeString[:len(routeString)-1]
	}

	fmt.Printf("%s ", routeString)

	if c.Skipped {
		fmt.Printf("Skipped\n")
	} else if !c.Result.Success {
		fmt.Printf("❌\n")
		fmt.Println(c.Result.Error)
	} else {
		fmt.Printf("✅\n")
	}
	fmt.Printf("=======\n")
}

// assertExpectations does expectations checking for the Case and returns an
// error if any stated expectations are not satisfied by actual results.
func (c Case) assertExpectations(result *oapi.ExecutorResult) error {
	// If the config has an expected resonse status, evaluate that.
	if c.Config.ExpectedRespStatus != 0 {
		if c.Config.ExpectedRespStatus != *result.StatusCode {
			return fmt.Errorf("Expected status code: %d\nReceived status code: %d", c.Config.ExpectedRespStatus, *result.StatusCode)
		}
	}

	// If the config has an expected response body, evaluate that.
	if c.Config.ExpectedRespBody != nil {
		// If the expected response is a simple string and not a JSON blob, do
		// that comparison here.
		if expectedString, ok := c.Config.ExpectedRespBody.(string); ok {
			actualString := reflect.ValueOf(result.Response).String()
			if strings.Compare(expectedString, actualString) != 0 {
				return fmt.Errorf("Expected response body:\n%s\n\nReceived response body:\n%s", expectedString, actualString)
			}
		} else {
			// If the expected response body is JSON, do that comparison that here.
			err := c.compareActualAndExpectedJson(result)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// compareActualAndExpectedJson unmarshals the expected response body into the
// appropriate Go type, only then to marshal it back out to JSON in a
// canonicalized form. The received JSON is already stored in its canonical
// data structure and is marshaled into canonicalized JSON bytes as well. The
// bytes are compared. An error is returned if the canonicalizes byte slices
// are not identical, or if there was an issue marshaling or unmarshaling any
// data.
func (c Case) compareActualAndExpectedJson(result *oapi.ExecutorResult) error {
	data, err := json.Marshal(c.Config.ExpectedRespBody)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &result.ResponseDS)
	if err != nil {
		return err
	}
	canonicalizedExpected, err := json.Marshal(result.ResponseDS)
	if err != nil {
		return err
	}

	canonicalizedActual, err := json.Marshal(result.Response)
	if err != nil {
		return err
	}

	if !bytes.Equal(canonicalizedExpected, canonicalizedActual) {
		return fmt.Errorf("Expected response body:\n%s\n\nReceived response body:\n%s", canonicalizedExpected, canonicalizedActual)
	}

	return nil
}

// awaitTargetIsHealthy blocks until the target server reports itself as being
// ready.
func (c Case) awaitTargetIsHealthy(ctx context.Context) error {
	for {
		if ctx.Err() != nil {
			return BadTargetError{Route: "/v1/node/health", Err: ctx.Err()}
		}

		httpdata, _ := c.OapiClient.NodeApi.GetHealth(ctx)
		switch {
		case httpdata != nil && (httpdata.StatusCode == 200 || httpdata.StatusCode == 206):
			return nil
		default:
			time.Sleep(time.Second)
			continue
		}
	}

}

type BadTargetError struct {
	Route string
	Err   error
}

func (e BadTargetError) Error() string {
	return fmt.Sprintf("BadTargetError: %s.\n\nDoes target implement %s? For information about the correct implementation of this required route, see https://ethereum.github.io/eth2.0-APIs/.", e.Err.Error(), e.Route)
}

// awaitTargetHasSlot blocks until the target server has synchronized the slot needed
// for the test case.
func (c Case) awaitTargetHasSlot(ctx context.Context) error {
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		headSlot, syncDistance, err := c.getHeadSlotAndSyncDistance(ctx)
		if err != nil {
			return err
		}

		currentSlot := headSlot - syncDistance

		switch {
		case currentSlot >= c.Config.AwaitSlot:
			return nil
		default:
			time.Sleep(time.Second)
			continue
		}
	}
}

// getHeadSlotAndSyncDistance is a convenience function that encapsulates some
// logic for awaitTargetHasSlot.
func (c Case) getHeadSlotAndSyncDistance(ctx context.Context) (int, int, error) {
	result, _, err := c.OapiClient.NodeApi.GetSyncingStatus(ctx)
	if err != nil {
		return 0, 0, BadTargetError{Route: "/v1/node/syncing", Err: err}
	}

	headSlot, err := strconv.ParseInt(result.Data.HeadSlot.(string), 10, 0)
	if err != nil {
		return 0, 0, err
	}
	syncDistance, err := strconv.ParseInt(result.Data.SyncDistance.(string), 10, 0)
	if err != nil {
		return 0, 0, err
	}

	return int(headSlot), int(syncDistance), nil
}
