// package target includes utilities for ascertaining the state of the target
// node, namely whether it is healthy and is ready to serve requests, and what
// slot it has synced to.
package target

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/INFURA/eth2-comply/pkg/oapi"
	"github.com/avast/retry-go"
)

type BadTargetError struct {
	Route string
	Err   error
}

func (e BadTargetError) Error() string {
	return fmt.Sprintf("BadTargetError: %s.\n\nDoes target implement %s? For information about the correct implementation of this required route, see https://ethereum.github.io/eth2.0-APIs/.", e.Err.Error(), e.Route)
}

// IsHealthy blocks until the target server reports itself as being ready.
func IsHealthy(ctx context.Context) error {
	for {
		if ctx.Err() != nil {
			return BadTargetError{Route: "/eth/v1/node/health", Err: ctx.Err()}
		}

		client := oapi.GetClient(ctx)
		httpdata, _ := client.NodeApi.GetHealth(ctx)
		switch {
		case httpdata != nil && (httpdata.StatusCode == 200 || httpdata.StatusCode == 206):
			return nil
		default:
			time.Sleep(time.Second)
			continue
		}
	}

}

type ClientMissingTargetSlotErr struct {
	Current int
	Target  int
}

func (e *ClientMissingTargetSlotErr) Error() string {
	return fmt.Sprintf("Target is at slot %d. Needs slot %d.", e.Current, e.Target)
}

// HasSlot blocks until the target server has synchronized the slot needed for
// the test case.
func HasSlot(ctx context.Context, awaitSlot int) error {
	var defaultRetryInterval time.Duration = time.Second
	var defaultRetryAttempts uint

	if deadline, ok := ctx.Deadline(); ok {
		timeToDeadline := time.Until(deadline)
		defaultRetryAttempts = uint((timeToDeadline.Seconds() / defaultRetryInterval.Seconds()))
	}

	if err := retry.Do(
		func() error {
			if ctx.Err() != nil {
				return retry.Unrecoverable(ctx.Err())
			}

			headSlot, _, err := getHeadSlotAndSyncDistance(ctx)
			if err != nil {
				return err
			}

			if headSlot >= awaitSlot {
				return nil
			}

			return &ClientMissingTargetSlotErr{Current: headSlot, Target: awaitSlot}
		},
		retry.Delay(time.Second),
		retry.Attempts(defaultRetryAttempts),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	); err != nil {
		return err
	}

	return nil
}

// getHeadSlotAndSyncDistance is a convenience function that encapsulates some
// logic for awaitTargetHasSlot.
func getHeadSlotAndSyncDistance(ctx context.Context) (int, int, error) {
	client := oapi.GetClient(ctx)
	result, _, err := client.NodeApi.GetSyncingStatus(ctx)
	if err != nil {
		return 0, 0, BadTargetError{Route: "/eth/v1/node/syncing", Err: err}
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
