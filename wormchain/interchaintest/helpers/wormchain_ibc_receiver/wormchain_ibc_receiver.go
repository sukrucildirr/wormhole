/* Code generated by github.com/srdtrk/go-codegen, DO NOT EDIT. */
package wormchain_ibc_receiver

type InstantiateMsg struct{}

type ExecuteMsg struct {
	// Submit one or more signed VAAs to update the on-chain state.  If processing any of the VAAs returns an error, the entire transaction is aborted and none of the VAAs are committed.
	SubmitUpdateChannelChain *ExecuteMsg_SubmitUpdateChannelChain `json:"submit_update_channel_chain,omitempty"`
}

// Contract queries
type QueryMsg struct {
	AllChannelChains *QueryMsg_AllChannelChains `json:"all_channel_chains,omitempty"`
	ChannelChain     *QueryMsg_ChannelChain     `json:"channel_chain,omitempty"`
}

type ExecuteMsg_SubmitUpdateChannelChain struct {
	// One or more VAAs to be submitted.  Each VAA should be encoded in the standard wormhole wire format.
	Vaas []Binary `json:"vaas"`
}

/*
Binary is a wrapper around Vec<u8> to add base64 de/serialization with serde. It also adds some helper methods to help encode inline.
This is only needed as serde-json-{core,wasm} has a horrible encoding for Vec<u8>. See also <https://github.com/CosmWasm/cosmwasm/blob/main/docs/MESSAGE_TYPES.md>.
*/
type Binary string

type QueryMsg_AllChannelChains struct{}

type QueryMsg_ChannelChain struct {
	ChannelId Binary `json:"channel_id"`
}

type ChannelChainResponse struct {
	ChainId int `json:"chain_id"`
}

type AllChannelChainsResponse struct {
	ChannelsChains []any `json:"channels_chains"`
}