// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

/// @custom:security-contact contract_security@leealabs.com
contract AgentRegistry {
    mapping(address => bool) private _agents;

    address private _owner;

    constructor(address initialOwner) {
        _owner = initialOwner;
    }
}
