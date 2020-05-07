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
	case "POST":
		return c.execPostOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/node/"):
		return c.execGetNodeOperation(ctx, route)
	case strings.Contains(route, "/beacon/"):
		return c.execGetBeaconOperation(ctx, route)

	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execPostOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/genesis"):
		return oapi.ExecGetBeaconGenesis(ctx, c.OapiClient)
	case strings.Contains(route, "/headers/"):
		blockId := uriTokens[4]
		return oapi.ExecGetBeaconHeader(ctx, c.OapiClient, blockId)
	case strings.Contains(route, "/headers"):
		return oapi.ExecGetBeaconHeaders(ctx, c.OapiClient, c.Config.QueryParams)
	case strings.Contains(route, "/blocks/"):
		blockId := uriTokens[4]
		switch {
		case strings.Contains(route, "/root"):
			return oapi.ExecGetBeaconBlockRoot(ctx, c.OapiClient, blockId)
		case strings.Contains(route, "/attestations"):
			return oapi.ExecGetBeaconBlockAttestations(ctx, c.OapiClient, blockId)
		default:
			return oapi.ExecGetBeaconBlock(ctx, c.OapiClient, blockId)
		}
	case strings.Contains(route, "/pool/"):
		return c.execGetBeaconPoolOperation(ctx, route)
	case strings.Contains(route, "/states/"):
		return c.execGetBeaconStatesOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetNodeOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/health"):
		return oapi.ExecGetNodeHealth(ctx, c.OapiClient)
	case strings.Contains(route, "/syncing"):
		return oapi.ExecGetNodeSyncing(ctx, c.OapiClient)
	case strings.Contains(route, "/version"):
		return oapi.ExecGetNodeVersion(ctx, c.OapiClient)
	case strings.Contains(route, "/peers/"):
		peerId := uriTokens[len(uriTokens)-1]
		return oapi.ExecGetNodePeer(ctx, c.OapiClient, peerId)
	case strings.Contains(route, "/peers"):
		return oapi.ExecGetNodePeers(ctx, c.OapiClient)
	case strings.Contains(route, "/identity"):
		return oapi.ExecGetNodeIdentity(ctx, c.OapiClient)

	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconPoolOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/attestations"):
		return oapi.ExecGetBeaconPoolAttestations(ctx, c.OapiClient, c.Config.QueryParams)
	case strings.Contains(route, "/attester_slashings"):
		return oapi.ExecGetBeaconPoolAttesterSlashings(ctx, c.OapiClient)
	case strings.Contains(route, "/proposer_slashings"):
		return oapi.ExecGetBeaconPoolProposerSlashings(ctx, c.OapiClient)
	case strings.Contains(route, "/voluntary_exits"):
		return oapi.ExecGetBeaconPoolVoluntaryExits(ctx, c.OapiClient)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconStatesOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")
	stateId := uriTokens[4]

	switch {
	case strings.Contains(route, "/committees/"):
		epoch := uriTokens[6]
		opts := &oapi.ExecGetBeaconStatesCommitteesOpts{
			StateId:     stateId,
			Epoch:       epoch,
			QueryParams: c.Config.QueryParams,
		}
		return oapi.ExecGetBeaconStatesCommittees(ctx, c.OapiClient, opts)
	case strings.Contains(route, "/finality_checkpoints"):
		return oapi.ExecGetBeaconStatesFinalityCheckpoints(ctx, c.OapiClient, stateId)
	case strings.Contains(route, "/fork"):
		return oapi.ExecGetBeaconStatesFork(ctx, c.OapiClient, stateId)
	case strings.Contains(route, "/root"):
		return oapi.ExecGetBeaconStatesRoot(ctx, c.OapiClient, stateId)
	case strings.Contains(route, "/validators/"):
		validatorId := uriTokens[6]
		opts := &oapi.ExecGetBeaconStatesValidatorsWithValidatorIdOpts{
			StateId:     stateId,
			ValidatorId: validatorId,
		}
		return oapi.ExecGetBeaconStatesValidator(ctx, c.OapiClient, opts)
	case strings.Contains(route, "/validators"):
		opts := &oapi.ExecGetBeaconStatesValidatorsOpts{
			StateId:     stateId,
			QueryParams: c.Config.QueryParams,
		}
		return oapi.ExecGetBeaconStatesValidators(ctx, c.OapiClient, opts)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}
