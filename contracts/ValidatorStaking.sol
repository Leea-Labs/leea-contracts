// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {LeeaGlobalParams} from "./GlobalParameters.sol";
import {ValidatorRegistry} from "./ValidatorRegistry.sol";

/// @custom:security-contact contract_security@leealabs.com
contract ValidatorStaking is Ownable {
    event StakeSuccess(
        address validator,
        address indexed staking,
        address indexed token,
        uint256 amount
    );

    mapping(address => uint256) private _stakes;
    ERC20 private _token;
    LeeaGlobalParams private _globalParams;
    ValidatorRegistry private _valRegistry;

    constructor(
        address initialOwner,
        ERC20 leeaToken,
        LeeaGlobalParams globalParams,
        ValidatorRegistry validatorReg
    ) Ownable(initialOwner) {
        _token = leeaToken;
        _globalParams = globalParams;
        _valRegistry = validatorReg;
    }

    function stake() public {
        require(
            _valRegistry.isValidator(msg.sender),
            "Validator is not registered"
        );
        uint256 minStake = _globalParams.getStake();
        require(
            _token.balanceOf(msg.sender) >= minStake,
            "Your balance should be greater or equal to required stake amount"
        );
        require(
            _token.approve(address(this), minStake),
            "Cant approve tokens to stake"
        );
        require(
            _token.transferFrom(msg.sender, address(this), minStake),
            "Cant transfer tokens to stake"
        );
        _stakes[msg.sender] += minStake;
        emit StakeSuccess(msg.sender, address(this), address(_token), minStake);
    }
}
