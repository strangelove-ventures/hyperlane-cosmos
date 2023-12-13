use crate::{
    msg::{IsmResponse, OwnerResponse},
    state::{ISM, OWNER},
};
use cosmwasm_std::{Deps, StdResult};

pub fn query_owner(deps: Deps) -> StdResult<OwnerResponse> {
    let owner = OWNER.load(deps.storage)?;
    Ok(OwnerResponse {
        address: owner.to_string(),
    })
}

pub fn query_ism(deps: Deps) -> StdResult<IsmResponse> {
    let ism = ISM.load(deps.storage)?;
    Ok(IsmResponse { ism_id: ism })
}