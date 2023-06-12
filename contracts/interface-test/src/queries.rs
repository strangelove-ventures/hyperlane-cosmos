use crate::{
    msg::OwnerResponse,
    state::OWNER,
};
use cosmwasm_std::{Deps, StdResult};

pub fn query_owner(deps: Deps) -> StdResult<OwnerResponse> {
    let owner = OWNER.load(deps.storage)?;
    Ok(OwnerResponse {
        address: owner.to_string(),
    })
}