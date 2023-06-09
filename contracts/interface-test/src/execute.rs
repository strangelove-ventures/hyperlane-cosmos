use crate::{
    error::ContractError,
    state::OWNER,
    hyperlane_bindings::HyperlaneMsg,
};
use cosmwasm_std::{Addr, Deps, DepsMut, MessageInfo, Response, CosmosMsg};

pub fn dispatch_msg(
    _deps: DepsMut,
    info: MessageInfo,
    destination: u32,
    recipient_address: String,
    message_body: String,
) -> Result<Response<HyperlaneMsg>, ContractError> {
    let message = HyperlaneMsg::MsgDispatch{
        sender: info.sender.clone().into_string(),
        destination_domain: destination.clone(),
        recipient_address: recipient_address.clone(),
        message_body: message_body.clone(),
    };

    let cosmos_message: CosmosMsg<HyperlaneMsg> =
        CosmosMsg::Custom(message);

    let msg_vec: Vec<u8> = prefix_hex::decode(message_body.clone()).unwrap();
    let msg_string = std::str::from_utf8(&msg_vec).unwrap();

    Ok(Response::new()
        .add_message(cosmos_message)
        .add_attribute("action", "dispatch_msg")
        .add_attribute("sender", info.sender)
        .add_attribute("destination_address", destination.to_string())
        .add_attribute("recipient_address", recipient_address)
        .add_attribute("message_body", msg_string))
}

pub fn process_msg(
    _deps: DepsMut,
    _info: MessageInfo,
    origin: u32,
    sender: String,
    msg: String,
) -> Result<Response<HyperlaneMsg>, ContractError> {
    Ok(Response::new()
    .add_attribute("action", "process_msg")
    .add_attribute("origin", origin.to_string())
    .add_attribute("sender", sender)
    .add_attribute("msg", msg))
}

pub fn change_contract_owner(
    deps: DepsMut,
    info: MessageInfo,
    new_owner: String,
) -> Result<Response<HyperlaneMsg>, ContractError> {
    // Only allow current contract owner to change owner
    check_is_contract_owner(deps.as_ref(), info.sender)?;

    // validate that new owner is a valid address
    let new_owner_addr = deps.api.addr_validate(&new_owner)?;

    // update the contract owner in the contract config
    OWNER.save(deps.storage, &new_owner_addr)?;

    // return OK
    Ok(Response::new()
        .add_attribute("action", "change_contract_owner")
        .add_attribute("new_owner", new_owner))
}

pub fn check_is_contract_owner(deps: Deps, sender: Addr) -> Result<(), ContractError> {
    let owner = OWNER.load(deps.storage)?;
    if owner != sender {
        Err(ContractError::Unauthorized {})
    } else {
        Ok(())
    }
}