// SPDX-License-Identifier: MIT
// Compatible with OpenZeppelin Contracts ^5.0.0
pragma solidity ^0.8.22;

import {GovernorUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/GovernorUpgradeable.sol";
import {GovernorCountingSimpleUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorCountingSimpleUpgradeable.sol";
import {GovernorSettingsUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorSettingsUpgradeable.sol";
import {GovernorStorageUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorStorageUpgradeable.sol";
import {GovernorTimelockControlUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorTimelockControlUpgradeable.sol";
import {GovernorVotesUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorVotesUpgradeable.sol";
import {GovernorVotesQuorumFractionUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/extensions/GovernorVotesQuorumFractionUpgradeable.sol";
import {Initializable} from "./lib/openzeppelin/contracts/proxy/utils/Initializable.sol";
import {TimelockControllerUpgradeable} from "./lib/openzeppelin/contracts-upgradable/governance/TimelockControllerUpgradeable.sol";

/// @custom:security-contact contract_security@leealabs.com
contract LeeaGovernance is Initializable, GovernorUpgradeable, GovernorSettingsUpgradeable, GovernorCountingSimpleUpgradeable, GovernorStorageUpgradeable, GovernorVotesUpgradeable, GovernorVotesQuorumFractionUpgradeable, GovernorTimelockControlUpgradeable {
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(IVotes _token, TimelockControllerUpgradeable _timelock)
        initializer public
    {
        __Governor_init("LeeaGovernance");
        __GovernorSettings_init(1 days, 1 weeks, 3e18);
        __GovernorCountingSimple_init();
        __GovernorStorage_init();
        __GovernorVotes_init(_token);
        __GovernorVotesQuorumFraction_init(4);
        __GovernorTimelockControl_init(_timelock);
    }

    // The following functions are overrides required by Solidity.

    function votingDelay()
        public
        view
        override(GovernorUpgradeable, GovernorSettingsUpgradeable)
        returns (uint256)
    {
        return super.votingDelay();
    }

    function votingPeriod()
        public
        view
        override(GovernorUpgradeable, GovernorSettingsUpgradeable)
        returns (uint256)
    {
        return super.votingPeriod();
    }

    function quorum(uint256 blockNumber)
        public
        view
        override(GovernorUpgradeable, GovernorVotesQuorumFractionUpgradeable)
        returns (uint256)
    {
        return super.quorum(blockNumber);
    }

    function state(uint256 proposalId)
        public
        view
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
        returns (ProposalState)
    {
        return super.state(proposalId);
    }

    function proposalNeedsQueuing(uint256 proposalId)
        public
        view
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
        returns (bool)
    {
        return super.proposalNeedsQueuing(proposalId);
    }

    function proposalThreshold()
        public
        view
        override(GovernorUpgradeable, GovernorSettingsUpgradeable)
        returns (uint256)
    {
        return super.proposalThreshold();
    }

    function _propose(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, string memory description, address proposer)
        internal
        override(GovernorUpgradeable, GovernorStorageUpgradeable)
        returns (uint256)
    {
        return super._propose(targets, values, calldatas, description, proposer);
    }

    function _queueOperations(uint256 proposalId, address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
        internal
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
        returns (uint48)
    {
        return super._queueOperations(proposalId, targets, values, calldatas, descriptionHash);
    }

    function _executeOperations(uint256 proposalId, address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
        internal
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
    {
        super._executeOperations(proposalId, targets, values, calldatas, descriptionHash);
    }

    function _cancel(address[] memory targets, uint256[] memory values, bytes[] memory calldatas, bytes32 descriptionHash)
        internal
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
        returns (uint256)
    {
        return super._cancel(targets, values, calldatas, descriptionHash);
    }

    function _executor()
        internal
        view
        override(GovernorUpgradeable, GovernorTimelockControlUpgradeable)
        returns (address)
    {
        return super._executor();
    }
}
