// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @custom:security-contact contract_security@leealabs.com
contract AgentRegistry is Ownable {
    event Registered(address pub, string name);
    struct Agent {
        string name;
        uint256 score;
        address wallet;
        uint index;
    }

    mapping(address => Agent) private _agentStructs;
    address[] private _agentIndex;

    constructor(address initialOwner) Ownable(initialOwner) {}

    function registerAgent(address agentAddress) public returns (uint index) {
        _agentIndex.push(agentAddress);
        _agentStructs[agentAddress].wallet = agentAddress;
        _agentStructs[agentAddress].index = _agentIndex.length - 1;
        return _agentIndex.length - 1;
    }

    function deleteAgent(address agentAddress) public returns (uint index) {
        uint rowToDelete = _agentStructs[agentAddress].index;
        address keyToMove = _agentIndex[_agentIndex.length - 1];
        _agentIndex[rowToDelete] = keyToMove;
        _agentStructs[keyToMove].index = rowToDelete;
        _agentIndex.pop();
        return rowToDelete;
    }
}
