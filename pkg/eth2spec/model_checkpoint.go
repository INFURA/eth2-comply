/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// Checkpoint The [`Checkpoint`](https://github.com/ethereum/eth2.0-specs/blob/v0.12.2/specs/phase0/beacon-chain.md#checkpoint
type Checkpoint struct {
	Epoch string `json:"epoch,omitempty"`
	Root string `json:"root,omitempty"`
}
