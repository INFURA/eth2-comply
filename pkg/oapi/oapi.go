// package oapi primarily provides executors for Ethereum 2.0 API operations.
package oapi

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/INFURA/eth2-comply/pkg/eth2spec"
	"github.com/antihax/optional"
)

// NewClient returns an OAPI client which can conduct operations against the
// provided target.
func NewClient(target url.URL) *eth2spec.APIClient {
	cfg := eth2spec.NewConfiguration()

	cfg.Host = target.Host

	client := eth2spec.NewAPIClient(cfg)

	return client
}

// ExecutorResult is the first element in the uniform tuple returned by all
// oapi executor functions.
type ExecutorResult struct {
	// Response is an instantiated OAPI data structure of the type returned
	// by the executor for a function's route.
	Response interface{}
	// ResponseDS (DS = Data Structure) is an empty OAPI data structure of
	// the type returned by the executor for a function's route. Properly
	// encoded expected responses can be unmarshaled into this struct. This
	// struct can be marshaled into a canonical JSON encoding to compare
	// against the actual response.
	ResponseDS interface{}
	// StatusCode is the HTTP status code returned for the operation executed
	// by the function. It can be used to compare against expected status
	// code results.
	StatusCode *int
}

func ExecGetBeaconGenesis(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {
	genesis, httpdata, err := client.BeaconApi.GetGenesis(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   genesis,
		ResponseDS: eth2spec.InlineResponse200{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconStatesFork(ctx context.Context, client *eth2spec.APIClient, stateId string) (*ExecutorResult, error) {
	fork, httpdata, err := client.BeaconApi.GetStateFork(ctx, stateId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   fork,
		ResponseDS: eth2spec.InlineResponse2002{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconStatesRoot(ctx context.Context, client *eth2spec.APIClient, stateId string) (*ExecutorResult, error) {
	root, httpdata, err := client.BeaconApi.GetStateRoot(ctx, stateId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   root,
		ResponseDS: eth2spec.InlineResponse2001{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconStatesFinalityCheckpoints(ctx context.Context, client *eth2spec.APIClient, stateId string) (*ExecutorResult, error) {
	finalityCheckpoint, httpdata, err := client.BeaconApi.GetStateFinalityCheckpoints(ctx, stateId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   finalityCheckpoint,
		ResponseDS: eth2spec.InlineResponse2003{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

type ExecGetBeaconStatesCommitteesOpts struct {
	StateId     string
	Epoch       string
	QueryParams map[string]string
}

func ExecGetBeaconStatesCommittees(ctx context.Context, client *eth2spec.APIClient, opts *ExecGetBeaconStatesCommitteesOpts) (*ExecutorResult, error) {

	getEpochCommitteesOpts := &eth2spec.GetEpochCommitteesOpts{}

	if len(opts.QueryParams["index"]) > 0 {
		getEpochCommitteesOpts.Index = optional.NewInterface(opts.QueryParams["index"])
	}
	if len(opts.QueryParams["slot"]) > 0 {
		getEpochCommitteesOpts.Slot = optional.NewInterface(opts.QueryParams["slot"])
	}

	committees, httpdata, err := client.BeaconApi.GetEpochCommittees(ctx, opts.StateId, opts.Epoch, getEpochCommitteesOpts)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   committees,
		ResponseDS: eth2spec.InlineResponse2006{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

type ExecGetBeaconStatesValidatorsOpts struct {
	StateId     string
	QueryParams map[string]string
}

func ExecGetBeaconStatesValidators(ctx context.Context, client *eth2spec.APIClient, opts *ExecGetBeaconStatesValidatorsOpts) (*ExecutorResult, error) {
	getStateValidatorsOpts := &eth2spec.GetStateValidatorsOpts{}

	if len(opts.QueryParams["id"]) > 0 {
		getStateValidatorsOpts.Id = optional.NewInterface(opts.QueryParams["id"])
	}
	if len(opts.QueryParams["status"]) > 0 {
		getStateValidatorsOpts.Status = optional.NewInterface(opts.QueryParams["status"])
	}

	validators, httpdata, err := client.BeaconApi.GetStateValidators(ctx, opts.StateId, getStateValidatorsOpts)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   validators,
		ResponseDS: eth2spec.InlineResponse2004{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

type ExecGetBeaconStatesValidatorsWithValidatorIdOpts struct {
	StateId     string
	ValidatorId string
}

func ExecGetBeaconStatesValidator(ctx context.Context, client *eth2spec.APIClient, opts *ExecGetBeaconStatesValidatorsWithValidatorIdOpts) (*ExecutorResult, error) {
	validator, httpdata, err := client.BeaconApi.GetStateValidator(ctx, opts.StateId, opts.ValidatorId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   validator,
		ResponseDS: eth2spec.InlineResponse2005{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconHeaders(ctx context.Context, client *eth2spec.APIClient, queryParams map[string]string) (*ExecutorResult, error) {
	getBlockHeaderOpts := &eth2spec.GetBlockHeadersOpts{}

	if len(queryParams["slot"]) > 0 {
		getBlockHeaderOpts.Slot = optional.NewInterface(queryParams["slot"])
	}
	if len(queryParams["parent_root"]) > 0 {
		getBlockHeaderOpts.ParentRoot = optional.NewInterface(queryParams["parent_root"])
	}

	headers, httpdata, err := client.BeaconApi.GetBlockHeaders(ctx, getBlockHeaderOpts)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   headers,
		ResponseDS: eth2spec.InlineResponse2007{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconHeader(ctx context.Context, client *eth2spec.APIClient, blockId string) (*ExecutorResult, error) {

	header, httpdata, err := client.BeaconApi.GetBlockHeader(ctx, blockId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   header,
		ResponseDS: eth2spec.InlineResponse2008{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconBlock(ctx context.Context, client *eth2spec.APIClient, blockId string) (*ExecutorResult, error) {
	block, httpdata, err := client.BeaconApi.GetBlock(ctx, blockId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   block,
		ResponseDS: eth2spec.InlineResponse2009{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconBlockRoot(ctx context.Context, client *eth2spec.APIClient, blockId string) (*ExecutorResult, error) {
	blockRoot, httpdata, err := client.BeaconApi.GetBlockRoot(ctx, blockId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   blockRoot,
		ResponseDS: eth2spec.InlineResponse20010{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconBlockAttestations(ctx context.Context, client *eth2spec.APIClient, blockId string) (*ExecutorResult, error) {
	blockAttestations, httpdata, err := client.BeaconApi.GetBlockAttestations(ctx, blockId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   blockAttestations,
		ResponseDS: eth2spec.InlineResponse20011{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconPoolAttestations(ctx context.Context, client *eth2spec.APIClient, queryParams map[string]string) (*ExecutorResult, error) {
	getPoolAttestationsOpts := &eth2spec.GetPoolAttestationsOpts{}

	if len(queryParams["slot"]) > 0 {
		getPoolAttestationsOpts.Slot = optional.NewString(queryParams["slot"])
	}
	if len(queryParams["committee_index"]) > 0 {
		getPoolAttestationsOpts.CommitteeIndex = optional.NewString(queryParams["committee_index"])
	}

	poolAttestations, httpdata, err := client.BeaconApi.GetPoolAttestations(ctx, getPoolAttestationsOpts)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   poolAttestations,
		ResponseDS: eth2spec.InlineResponse20011{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconPoolAttesterSlashings(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	attesterSlashings, httpdata, err := client.BeaconApi.GetPoolAttesterSlashings(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   attesterSlashings,
		ResponseDS: eth2spec.InlineResponse20012{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconPoolProposerSlashings(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	proposerSlashings, httpdata, err := client.BeaconApi.GetPoolProposerSlashings(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   proposerSlashings,
		ResponseDS: eth2spec.InlineResponse20013{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetBeaconPoolVoluntaryExits(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	voluntaryExits, httpdata, err := client.BeaconApi.GetPoolVoluntaryExits(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   voluntaryExits,
		ResponseDS: eth2spec.InlineResponse20014{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodeHealth(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	httpdata, err := client.NodeApi.GetHealth(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   nil,
		ResponseDS: nil,
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodeSyncing(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	syncing, httpdata, err := client.NodeApi.GetSyncingStatus(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   syncing,
		ResponseDS: eth2spec.InlineResponse20019{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodeVersion(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	version, httpdata, err := client.NodeApi.GetNodeVersion(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   version,
		ResponseDS: eth2spec.InlineResponse20018{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodeIdentity(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	identity, httpdata, err := client.NodeApi.GetNetworkIdentity(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   identity,
		ResponseDS: eth2spec.InlineResponse20015{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodePeers(ctx context.Context, client *eth2spec.APIClient) (*ExecutorResult, error) {

	peers, httpdata, err := client.NodeApi.GetPeers(ctx)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   peers,
		ResponseDS: eth2spec.InlineResponse20016{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecGetNodePeer(ctx context.Context, client *eth2spec.APIClient, peerId string) (*ExecutorResult, error) {

	peer, httpdata, err := client.NodeApi.GetPeer(ctx, peerId)
	if err != nil {
		return nil, err
	}

	result := &ExecutorResult{
		Response:   peer,
		ResponseDS: eth2spec.InlineResponse20017{},
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}

func ExecPostBeaconPoolVoluntaryExits(ctx context.Context, client *eth2spec.APIClient, requestBody interface{}) (*ExecutorResult, error) {
	data, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	voluntaryExit := &eth2spec.InlineObject3{}
	err = json.Unmarshal(data, voluntaryExit)
	if err != nil {
		return nil, err
	}

	httpdata, err := client.BeaconApi.SubmitPoolVoluntaryExit(ctx, *voluntaryExit)

	result := &ExecutorResult{
		Response:   nil,
		ResponseDS: nil,
		StatusCode: &httpdata.StatusCode,
	}

	return result, nil
}
