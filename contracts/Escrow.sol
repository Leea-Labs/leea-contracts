// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @custom:security-contact contract_security@leealabs.com
contract Escrow is Ownable, ReentrancyGuard {
    event Deposited(
        address user,
        address indexed escrow,
        address indexed token,
        uint256 amount
    );
    event Withdrawn(
        address user,
        address indexed escrow,
        address indexed token,
        uint256 amount
    );
    event FeePaid(
        address user,
        address agent,
        address indexed escrow,
        address indexed token,
        uint256 amount
    );

    mapping(address => uint256) private escrowBalance;

    ERC20 private _token;
    uint256 private _fee;

    constructor(ERC20 leeaToken, address initialOwner) Ownable(initialOwner) {
        _token = leeaToken;
    }

    function deposit(uint256 _amount) public {
        require(_amount > 0, "You need to deposit at least some tokens");
        require(
            _token.balanceOf(msg.sender) > _amount,
            "You need to have at least amount of tokens"
        );
        require(
            _token.approve(address(this), _amount),
            "Cant approve tokens to escrow"
        );
        require(
            _token.transferFrom(msg.sender, address(this), _amount),
            "Cant transfer tokens to escrow"
        );
        escrowBalance[msg.sender] += _amount;
        emit Deposited(msg.sender, address(this), address(_token), _amount);
    }

    function withdrawFullAmount(address _user) public nonReentrant {
        require(msg.sender == owner(), "Only owner is allowed");
        require(escrowBalance[_user] > 0, "Balance is zero");
        uint256 _currentBalance = escrowBalance[_user];
        require(
            _token.transferFrom(address(this), _user, _currentBalance),
            "Cant transfer tokens to user"
        );
        emit Withdrawn(_user, address(this), address(_token), _currentBalance);
    }

    function payFee(address _user, address _agent) public nonReentrant {
        require(msg.sender == owner(), "Only owner is allowed");
        require(escrowBalance[_user] >= _fee, "Balance less than fee");
        require(
            _token.transferFrom(address(this), _agent, _fee),
            "Cant transfer tokens to escrow"
        );
        escrowBalance[msg.sender] -= _fee;
        emit FeePaid(_user, _agent, address(this), address(_token), _fee);
    }
}
