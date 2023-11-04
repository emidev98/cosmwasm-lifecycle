use std::ops::Add;

use cosmwasm_std::{DepsMut, Env, MessageInfo,entry_point, Response, StdError};
use models::{SudoMsg, InstantiateMsg};

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn instantiate(
    _deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: InstantiateMsg,
) -> Result<Response, StdError> {
    Ok(Response::new())
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn sudo(
    _deps: DepsMut, 
    env: Env, 
    msg: SudoMsg
) -> Result<Response, StdError> {
    match msg {
        SudoMsg::SudoBeginBlock{} => begin_block(env),
        SudoMsg::SudoEndBlock{} => end_block(env),
    }
}

fn begin_block(env: Env) -> Result<Response, StdError> {
    let msg = String::from("begin_block error at height ").add(&env.block.height.to_string());
    Err(StdError::generic_err(msg))
}

fn end_block(env: Env) -> Result<Response, StdError> {
    let msg = String::from("end_block error at height ").add(&env.block.height.to_string());
    Err(StdError::generic_err(msg))
}