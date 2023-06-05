package helpers

// Go based data types for querying on the contract.

// TODO: Auto generate in the future from Rust types -> Go types?
// Execute types are not needed here. We just use strings. Could add though in the future and to_string it

// EntryPoint
type QueryMsg struct {
	Owner *struct{} `json:"owner,omitempty"`
}

type QueryRsp struct {
	Data *QueryRspObj `json:"data,omitempty"`
}

type QueryRspObj struct {
	Address string `json:"address,omitempty"`
}

type ExecuteMsg struct {
	DispatchMsg         *DispatchMsg         `json:"dispatch_msg,omitempty"`
	ProcessMsg          *ProcessMsg          `json:"process_msg,omitempty"`
	ChangeContractOwner *ChangeContractOwner `json:"change_contract_owner,omitempty"`
}

type ExecuteRsp struct{}

type DispatchMsg struct {
	DestinationAddr uint32 `json:"destination_domain"`
	RecipientAddr   string `json:"recipient_address"`
	MessageBody     string `json:"message_body"`
}

type ProcessMsg struct {
	Msg string `json:"msg"`
}

type ChangeContractOwner struct {
	NewOwner string `json:"new_owner"`
}
