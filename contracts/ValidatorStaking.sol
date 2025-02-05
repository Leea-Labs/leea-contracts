// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

/// @custom:security-contact contract_security@leealabs.com
contract Staking {
    mapping(address => uint256) private _amount;

    address private _owner;

    constructor(address initialOwner) {
        _owner = initialOwner;
    }
}
