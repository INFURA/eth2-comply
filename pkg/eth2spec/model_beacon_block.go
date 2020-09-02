/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// BeaconBlock The [`BeaconBlock`](https://github.com/ethereum/eth2.0-specs/blob/v0.12.2/specs/phase0/beacon-chain.md#beaconblock) object from the Eth2.0 spec.
type BeaconBlock struct {
	Slot string `json:"slot,omitempty"`
	ProposerIndex string `json:"proposer_index,omitempty"`
	ParentRoot string `json:"parent_root,omitempty"`
	StateRoot string `json:"state_root,omitempty"`
	Body BeaconBlockAllOf1Body `json:"body,omitempty"`
}
