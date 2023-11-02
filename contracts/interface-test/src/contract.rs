use crate::{
    error::ContractError,
    msg::{ExecuteMsg, InstantiateMsg, MigrateMsg, QueryMsg},
    state::{
        OWNER, ISM,
    },
    execute, queries,
    hyperlane_bindings::HyperlaneMsg,
};
use cosmwasm_std::{
    entry_point,
    to_binary, Binary, Deps, DepsMut, Env, Event, MessageInfo, Response, StdResult,
};

//#[entry_point]
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    OWNER.save(deps.storage, &info.sender)?;
    ISM.save(deps.storage, &0)?;

    Ok(Response::new()
        .add_event(Event::new("hyperlane_init").add_attribute("attr", "value"))
        .add_attribute("action", "instantiate")
        .add_attribute("owner", info.sender))
}

/// Allow contract to be able to migrate if admin is set.
/// This provides option for migration, if admin is not set, this functionality will be disabled
//#[entry_point]
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn migrate(_deps: DepsMut, _env: Env, _msg: MigrateMsg) -> Result<Response, ContractError> {
    Ok(Response::new().add_attribute("action", "migrate"))
}

//#[entry_point]
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn execute(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response<HyperlaneMsg>, ContractError> {
    match msg {
        // Admin functions
        ExecuteMsg::DispatchMsg { destination_domain, recipient_address, message_body } => {
            execute::dispatch_msg(deps, info, destination_domain, recipient_address, message_body)
        }
        ExecuteMsg::ProcessMsg { origin, sender, msg } => {
            execute::process_msg(deps, info, origin, sender, msg)
        }
        ExecuteMsg::ChangeContractOwner { new_owner } => {
            execute::change_contract_owner(deps, info, new_owner)
        }
        ExecuteMsg::SetIsmId { ism_id } => {
            execute::set_ism(deps, info, ism_id)
        }
    }
}

//#[entry_point]
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Owner {} => to_binary(&queries::query_owner(deps)?),
        QueryMsg::Ism {} => to_binary(&queries::query_ism(deps)?),
    }
}
