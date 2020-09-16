package testcases

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/INFURA/eth2-comply/pkg/oapi"
)

type UnimplementedOperationError struct {
	method string
	route  string
}

func (e UnimplementedOperationError) Error() string {
	return fmt.Sprintf("Tests for the operation %s %s are not supported.", e.method, e.route)
}

// execOperation is a big if-else tree that uses the CaseConfig method and
// route to determine the appropriate OAPI executor to use for actually
// executing the operation under test.
//
// For routes that use path params, it is the responsibility of the case for
// that route in this function to extract the path param and pass it along to
// the constituent oapi executor function.
func (c Case) execOperation(ctx context.Context) (*oapi.ExecutorResult, error) {
	_route, err := url.Parse(c.Config.Route)
	if err != nil {
		return nil, err
	}
	route := _route.RequestURI()

	switch c.Config.Method {
	case "GET":
		return c.execGetOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/validator/"):
		return c.execGetValidatorOperation(ctx, route)
	case strings.Contains(route, "/node/"):
		return c.execGetNodeOperation(ctx, route)
	case strings.Contains(route, "/debug/"):
		return c.execGetDebugOperation(ctx, route)
	case strings.Contains(route, "/beacon/"):
		return c.execGetBeaconOperation(ctx, route)
	case strings.Contains(route, "/config/"):
		return c.execGetConfigOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/genesis"):
		return oapi.ExecGetBeaconGenesis(ctx)
	case strings.Contains(route, "/headers/"):
		blockId := uriTokens[5]
		return oapi.ExecGetBeaconHeader(ctx, blockId)
	case strings.Contains(route, "/headers"):
		return oapi.ExecGetBeaconHeaders(ctx, c.Config.QueryParams)
	case strings.Contains(route, "/blocks/"):
		blockId := uriTokens[5]
		switch {
		case strings.Contains(route, "/root"):
			return oapi.ExecGetBeaconBlockRoot(ctx, blockId)
		case strings.Contains(route, "/attestations"):
			return oapi.ExecGetBeaconBlockAttestations(ctx, blockId)
		default:
			return oapi.ExecGetBeaconBlock(ctx, blockId)
		}
	case strings.Contains(route, "/states/"):
		return c.execGetBeaconStatesOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetNodeOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/health"):
		return oapi.ExecGetNodeHealth(ctx)
	case strings.Contains(route, "/syncing"):
		return oapi.ExecGetNodeSyncing(ctx)
	case strings.Contains(route, "/version"):
		return oapi.ExecGetNodeVersion(ctx)
	case strings.Contains(route, "/peers/"):
		peerId := uriTokens[len(uriTokens)-1]
		return oapi.ExecGetNodePeer(ctx, peerId)
	case strings.Contains(route, "/peers"):
		return oapi.ExecGetNodePeers(ctx)
	case strings.Contains(route, "/identity"):
		return oapi.ExecGetNodeIdentity(ctx)

	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconStatesOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")
	stateId := uriTokens[5]

	switch {
	case strings.Contains(route, "/committees/"):
		epoch := uriTokens[7]
		opts := &oapi.ExecGetBeaconStatesCommitteesOpts{
			StateId:     stateId,
			Epoch:       epoch,
			QueryParams: c.Config.QueryParams,
		}
		return oapi.ExecGetBeaconStatesCommittees(ctx, opts)
	case strings.Contains(route, "/finality_checkpoints"):
		return oapi.ExecGetBeaconStatesFinalityCheckpoints(ctx, stateId)
	case strings.Contains(route, "/fork"):
		return oapi.ExecGetBeaconStatesFork(ctx, stateId)
	case strings.Contains(route, "/root"):
		return oapi.ExecGetBeaconStatesRoot(ctx, stateId)
	case strings.Contains(route, "/validators/"):
		validatorId := uriTokens[7]
		return oapi.ExecGetBeaconStatesValidator(ctx, stateId, validatorId)
	case strings.Contains(route, "/validators"):
		return oapi.ExecGetBeaconStatesValidators(ctx, stateId, c.Config.QueryParams)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetDebugOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/states/"):
		stateId := uriTokens[6]
		return oapi.ExecGetDebugBeaconStates(ctx, stateId)
	case strings.Contains(route, "/heads"):
		return oapi.ExecGetDebugBeaconHeads(ctx)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetConfigOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/fork_schedule"):
		return oapi.ExecGetConfigForkSchedule(ctx)
	case strings.Contains(route, "/spec"):
		return oapi.ExecGetConfigSpec(ctx)
	case strings.Contains(route, "/deposit_contract"):
		return oapi.ExecGetConfigDepositContract(ctx)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetValidatorOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/duties/"):
		switch {
		case strings.Contains(route, "/attester/"):
			opts := oapi.ExecGetValidatorDutiesAttesterOpts{
				Epoch:       uriTokens[6],
				QueryParams: c.Config.QueryParams,
			}
			return oapi.ExecGetValidatorDutiesAttester(ctx, opts)
		case strings.Contains(route, "/proposer/"):
			return oapi.ExecGetValidatorDutiesProposer(ctx, uriTokens[6])
		}
	case strings.Contains(route, "/blocks/"):
		opts := oapi.ExecGetValidatorBlocksOpts{
			Slot:        uriTokens[5],
			QueryParams: c.Config.QueryParams,
		}
		return oapi.ExecGetValidatorBlocks(ctx, opts)
	case strings.Contains(route, "/attestation_data"):
		return oapi.ExecGetValidatorAttestationData(ctx, c.Config.QueryParams)
	case strings.Contains(route, "/aggregate_attestation"):
		return oapi.ExecGetValidatorAggregateAttestation(ctx, c.Config.QueryParams)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}
