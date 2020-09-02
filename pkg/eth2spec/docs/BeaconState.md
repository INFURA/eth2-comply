# BeaconState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenesisTime** | **string** |  | [optional] 
**GenesisValidatorsRoot** | **string** |  | [optional] 
**Slot** | **string** |  | [optional] 
**Fork** | [**GetStateForkResponseData**](GetStateForkResponse_data.md) |  | [optional] 
**LatestBlockHeader** | **map[string]interface{}** | The [&#x60;BeaconBlockHeader&#x60;](https://github.com/ethereum/eth2.0-specs/blob/v0.12.2/specs/phase0/beacon-chain.md#beaconblockheader) object from the Eth2.0 spec. | [optional] 
**BlockRoots** | **[]string** |  | [optional] 
**StateRoots** | **[]string** |  | [optional] 
**HistoricalRoots** | **[]string** |  | [optional] 
**Eth1Data** | [**BeaconStateEth1Data**](BeaconState_eth1_data.md) |  | [optional] 
**Eth1DataVotes** | **[]map[string]interface{}** |  | [optional] 
**Eth1DepositIndex** | **string** |  | [optional] 
**Validators** | **[]map[string]interface{}** |  | [optional] 
**Balances** | **[]string** | Validator balances in gwei | [optional] 
**RandaoMixes** | **[]string** |  | [optional] 
**Slashings** | **[]string** | Per-epoch sums of slashed effective balances | [optional] 
**PreviousEpochAttestations** | **[]map[string]interface{}** |  | [optional] 
**CurrentEpochAttestations** | **[]map[string]interface{}** |  | [optional] 
**JustificationBits** | **string** | Bit set for every recent justified epoch | [optional] 
**PreviousJustifiedCheckpoint** | [**GetStateFinalityCheckpointsResponseDataPreviousJustified**](GetStateFinalityCheckpointsResponse_data_previous_justified.md) |  | [optional] 
**CurrentJustifiedCheckpoint** | [**GetStateFinalityCheckpointsResponseDataPreviousJustified**](GetStateFinalityCheckpointsResponse_data_previous_justified.md) |  | [optional] 
**FinalizedCheckpoint** | [**GetStateFinalityCheckpointsResponseDataPreviousJustified**](GetStateFinalityCheckpointsResponse_data_previous_justified.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


