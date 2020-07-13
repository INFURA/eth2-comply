// package testcases defines the Case data structure, and a public function for
// instantiating slices of Cases from JSON test specifications in directory
// trees on the filesystem.
package testcases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/INFURA/eth2-comply/pkg/eth2spec"
	"github.com/INFURA/eth2-comply/pkg/oapi"
	"github.com/INFURA/eth2-comply/pkg/target"
)

// Case is an executable test case. The Config property can be accessed to get
// information about the case scenario.
type Case struct {
	Config  CaseConfig
	Result  Result
	Skipped bool
	Done    chan struct{}
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

type OapiError struct {
	Err            error
	ServerResponse []byte
}

func (e OapiError) Error() string {
	return fmt.Sprintf("OpenAPI client error!\nError: %s\nServer message: %s", e.Err.Error(), string(e.ServerResponse))
}

// NewCase instantiates and returns a Case struct.
func NewCase(config CaseConfig) *Case {
	c := &Case{}

	c.Config = config
	c.Result = Result{}
	c.Done = make(chan struct{})

	return c
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

	// If a test specifies an await slot, wait for the node to sync that slot.
	if c.Config.AwaitSlot > 0 {
		if err := target.HasSlot(ctx, c.Config.AwaitSlot); err != nil {
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
		c.setFailure(err)
		return
	}

	c.Result.Success = true
}

// ResultsPretty returns human-readable test results output suitable for
// printing to a CLI.
func (c Case) ResultsPretty() string {
	// Build up a route string with the query params appended to the end
	routeString := fmt.Sprintf("%s %s", c.Config.Method, c.Config.Route)
	if len(c.Config.QueryParams) > 0 {
		routeString = fmt.Sprintf("%s%s", routeString, "?")
		for i, x := range c.Config.QueryParams {
			routeString = fmt.Sprintf("%s%s=%s&", routeString, i, x)
		}
		// Remove the trailing ampersand
		routeString = routeString[:len(routeString)-1]
	}

	var resultString string
	if c.Skipped {
		resultString = fmt.Sprintf("%s skipped\n", routeString)
	} else if !c.Result.Success {
		resultString = fmt.Sprintf("%s ❌\n%s", routeString, c.Result.Error.Error())
	} else {
		resultString = fmt.Sprintf("%s ✅\n", routeString)
	}

	return resultString
}

// setFailure marks a test case as having failed and records a corresponding
// error.
func (c *Case) setFailure(err error) {
	c.Result.Success = false
	c.Result.Error = err
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
			// If the expected response body is JSON, do that comparison that
			// here.
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
