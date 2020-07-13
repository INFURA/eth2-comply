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
	switch {
	case strings.Contains(route, "/beacon/"):
		return c.execPostBeaconOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execGetBeaconOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	uriTokens := strings.Split(route, "/")

	switch {
	case strings.Contains(route, "/genesis"):
		return oapi.ExecGetBeaconGenesis(ctx)
	case strings.Contains(route, "/headers/"):
		blockId := uriTokens[4]
		return oapi.ExecGetBeaconHeader(ctx, blockId)
	case strings.Contains(route, "/headers"):
		return oapi.ExecGetBeaconHeaders(ctx, c.Config.QueryParams)
	case strings.Contains(route, "/blocks/"):
		blockId := uriTokens[4]
		switch {
		case strings.Contains(route, "/root"):
			return oapi.ExecGetBeaconBlockRoot(ctx, blockId)
		case strings.Contains(route, "/attestations"):
			return oapi.ExecGetBeaconBlockAttestations(ctx, blockId)
		default:
			return oapi.ExecGetBeaconBlock(ctx, blockId)
		}
	case strings.Contains(route, "/pool/"):
		return c.execGetBeaconPoolOperation(ctx, route)
	case strings.Contains(route, "/states/"):
		return c.execGetBeaconStatesOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execPostBeaconOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/pool/"):
		return c.execPostBeaconPoolOperation(ctx, route)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}

func (c Case) execPostBeaconPoolOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/voluntary_exits"):
		return oapi.ExecPostBeaconPoolVoluntaryExits(ctx, c.Config.ReqBody)
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

func (c Case) execGetBeaconPoolOperation(ctx context.Context, route string) (*oapi.ExecutorResult, error) {
	switch {
	case strings.Contains(route, "/attestations"):
		return oapi.ExecGetBeaconPoolAttestations(ctx, c.Config.QueryParams)
	case strings.Contains(route, "/attester_slashings"):
		return oapi.ExecGetBeaconPoolAttesterSlashings(ctx)
	case strings.Contains(route, "/proposer_slashings"):
		return oapi.ExecGetBeaconPoolProposerSlashings(ctx)
	case strings.Contains(route, "/voluntary_exits"):
		return oapi.ExecGetBeaconPoolVoluntaryExits(ctx)
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
		return oapi.ExecGetBeaconStatesCommittees(ctx, opts)
	case strings.Contains(route, "/finality_checkpoints"):
		return oapi.ExecGetBeaconStatesFinalityCheckpoints(ctx, stateId)
	case strings.Contains(route, "/fork"):
		return oapi.ExecGetBeaconStatesFork(ctx, stateId)
	case strings.Contains(route, "/root"):
		return oapi.ExecGetBeaconStatesRoot(ctx, stateId)
	case strings.Contains(route, "/validators/"):
		validatorId := uriTokens[6]
		return oapi.ExecGetBeaconStatesValidator(ctx, stateId, validatorId)
	case strings.Contains(route, "/validators"):
		return oapi.ExecGetBeaconStatesValidators(ctx, stateId, c.Config.QueryParams)
	}

	return nil, UnimplementedOperationError{method: c.Config.Method, route: route}
}
