use cosmwasm_schema::cw_serde;
use cosmwasm_std::CustomMsg;

#[cw_serde]
pub enum HyperlaneMsg {
    MsgDispatch {
        sender: String,
        destination_domain: u32,
        recipient_address: String,
        message_body: String,
    },
}

impl CustomMsg for HyperlaneMsg {}