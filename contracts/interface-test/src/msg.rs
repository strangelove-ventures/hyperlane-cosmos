use cosmwasm_schema::{cw_serde, QueryResponses};

#[cw_serde]
pub struct InstantiateMsg {}

#[cw_serde]
pub struct MigrateMsg {}

#[cw_serde]
pub enum ExecuteMsg {
    DispatchMsg {
        destination_domain: u32,
        recipient_address: String,
        message_body: String,
    },
    ProcessMsg {
        origin: u32,
        sender: String,
        msg: String,
    },
    ChangeContractOwner {
        new_owner: String,
    },
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    /// Owner returns the owner of the contract. Response: OwnerResponse
    #[returns(OwnerResponse)]
    Owner {},
}

#[cw_serde]
pub struct OwnerResponse {
    pub address: String,
}
