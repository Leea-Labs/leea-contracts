// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LeeaGovernanceMetaData contains all meta data concerning the LeeaGovernance contract.
var LeeaGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDetails\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"proposalDetailsAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b506040516200559a3803806200559a833981016040819052620000359162000610565b806004836040518060400160405280600e81526020016d4c656561476f7665726e616e636560901b81525080620000716200016360201b60201c565b6200007e8260006200017e565b610120526200008f8160016200017e565b61014052815160208084019190912060e052815190820120610100524660a0526200011d60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60805250503060c0526003620001348282620006f6565b50506001600160a01b0316610160526200014e81620001b7565b506200015a816200025b565b50505062000884565b6040805180820190915260018152603160f81b602082015290565b60006020835110156200019e576200019683620002c4565b9050620001b1565b81620001ab8482620006f6565b5060ff90505b92915050565b606480821115620001ea5760405163243e544560e01b815260048101839052602481018290526044015b60405180910390fd5b6000620001f662000307565b90506200021b6200020662000323565b6200021185620003a5565b600a9190620003df565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b600080829050601f81511115620002f2578260405163305a27a960e01b8152600401620001e19190620007c2565b8051620002ff8262000813565b179392505050565b600062000315600a620003fc565b6001600160d01b0316905090565b6000620003306101605190565b6001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200038c575060408051601f3d908101601f19168201909252620003899181019062000838565b60015b620003a0576200039b6200044c565b905090565b919050565b60006001600160d01b03821115620003db576040516306dfcc6560e41b815260d0600482015260248101839052604401620001e1565b5090565b600080620003ef85858562000459565b915091505b935093915050565b80546000908015620004425762000428836200041a60018462000862565b600091825260209091200190565b54660100000000000090046001600160d01b031662000445565b60005b9392505050565b60006200039b43620005c2565b825460009081908015620005635760006200047b876200041a60018562000862565b805490915065ffffffffffff80821691660100000000000090046001600160d01b0316908816821115620004c257604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff16036200050057825465ffffffffffff1666010000000000006001600160d01b0389160217835562000554565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f5560008f815291909120945191519092166601000000000000029216919091179101555b9450859350620003f492505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a5560008a8152918220955192519093166601000000000000029190931617920191909155905081620003f4565b600065ffffffffffff821115620003db576040516306dfcc6560e41b81526030600482015260248101839052604401620001e1565b6001600160a01b03811681146200060d57600080fd5b50565b600080604083850312156200062457600080fd5b82516200063181620005f7565b60208401519092506200064481620005f7565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200067a57607f821691505b6020821081036200069b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620006f1576000816000526020600020601f850160051c81016020861015620006cc5750805b601f850160051c820191505b81811015620006ed57828155600101620006d8565b5050505b505050565b81516001600160401b038111156200071257620007126200064f565b6200072a8162000723845462000665565b84620006a1565b602080601f831160018114620007625760008415620007495750858301515b600019600386901b1c1916600185901b178555620006ed565b600085815260208120601f198616915b82811015620007935788860151825594840194600190910190840162000772565b5085821015620007b25787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020808352835180602085015260005b81811015620007f257858101830151858201604001528201620007d4565b506000604082860101526040601f19601f8301168501019250505092915050565b805160208083015191908110156200069b5760001960209190910360031b1b16919050565b6000602082840312156200084b57600080fd5b815165ffffffffffff811681146200044557600080fd5b81810381811115620001b157634e487b7160e01b600052601160045260246000fd5b60805160a05160c05160e05161010051610120516101405161016051614c946200090660003960008181610a79015281816112cc0152818161179d0152818161243301526125ae015260006123fe015260006123d101526000612bf101526000612bc901526000612b2401526000612b4e01526000612b780152614c946000f3fe6080604052600436106102e85760003560e01c80637d5e81e211610190578063c01f9e37116100dc578063ddf0b00911610095578063f23a6e611161006f578063f23a6e6114610a2a578063f8ce560a14610a4a578063fc0c546a14610a6a578063fe0d94c114610a9d57600080fd5b8063ddf0b009146109b6578063deaaa7cc146109d6578063eb9019d414610a0a57600080fd5b8063c01f9e37146108ea578063c28bc2fa1461090a578063c59057e41461091d578063d33219b41461093d578063da35c6641461095b578063dd4e2ba51461097057600080fd5b80639a802a6d11610149578063a9a9529411610123578063a9a952941461085e578063ab58fb8e1461087e578063b58131b0146108b6578063bc197c81146108ca57600080fd5b80639a802a6d14610809578063a7713a7014610829578063a890c9101461083e57600080fd5b80637d5e81e21461072b5780637ecebe001461074b57806384b0196e146107815780638ff262e3146107a957806391ddadf4146107c957806397c3d334146107f557600080fd5b80633932abb11161024f578063544ffc9c116102085780635b8d0e0d116101e25780635b8d0e0d146106ab5780635f398a14146106cb57806360c4247f146106eb5780637b3c71d31461070b57600080fd5b8063544ffc9c1461060c57806354fd4d5014610661578063567813881461068b57600080fd5b80633932abb11461052a5780633e4f49e61461054057806340e58ee51461056d578063438596321461058d578063452115d6146105d75780634bf5d7e9146105f757600080fd5b8063160cbed7116102a1578063160cbed71461044257806316e9eaec146104625780632656227d146104925780632d63f693146104a55780632e82db94146104c55780632fe3e261146104f657600080fd5b806301ffc9a71461032457806302a251a31461035957806306f3f9e61461037957806306fdde0314610399578063143489d0146103bb578063150b7a021461040957600080fd5b3661031f57306102f6610ab0565b6001600160a01b03161461031d57604051637485328f60e11b815260040160405180910390fd5b005b600080fd5b34801561033057600080fd5b5061034461033f366004613ae6565b610ac9565b60405190151581526020015b60405180910390f35b34801561036557600080fd5b5062093a805b604051908152602001610350565b34801561038557600080fd5b5061031d610394366004613b10565b610b1b565b3480156103a557600080fd5b506103ae610b2f565b6040516103509190613b79565b3480156103c757600080fd5b506103f16103d6366004613b10565b6000908152600460205260409020546001600160a01b031690565b6040516001600160a01b039091168152602001610350565b34801561041557600080fd5b50610429610424366004613c6c565b610bc1565b6040516001600160e01b03199091168152602001610350565b34801561044e57600080fd5b5061036b61045d366004613e4b565b610c04565b34801561046e57600080fd5b5061048261047d366004613b10565b610cd3565b6040516103509493929190613faa565b61036b6104a0366004613e4b565b610edd565b3480156104b157600080fd5b5061036b6104c0366004613b10565b611049565b3480156104d157600080fd5b506104e56104e0366004613b10565b61106a565b604051610350959493929190613ff5565b34801561050257600080fd5b5061036b7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b34801561053657600080fd5b506201518061036b565b34801561054c57600080fd5b5061056061055b366004613b10565b6110ad565b604051610350919061407f565b34801561057957600080fd5b5061031d610588366004613b10565b6110b8565b34801561059957600080fd5b506103446105a836600461408d565b60008281526007602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b3480156105e357600080fd5b5061036b6105f2366004613e4b565b611259565b34801561060357600080fd5b506103ae6112c8565b34801561061857600080fd5b50610646610627366004613b10565b6000908152600760205260409020805460018201546002909201549092565b60408051938452602084019290925290820152606001610350565b34801561066d57600080fd5b506040805180820190915260018152603160f81b60208201526103ae565b34801561069757600080fd5b5061036b6106a63660046140ce565b61138a565b3480156106b757600080fd5b5061036b6106c6366004614142565b6113b3565b3480156106d757600080fd5b5061036b6106e63660046141fc565b611512565b3480156106f757600080fd5b5061036b610706366004613b10565b611567565b34801561071757600080fd5b5061036b61072636600461427f565b6115f5565b34801561073757600080fd5b5061036b6107463660046142d8565b61163d565b34801561075757600080fd5b5061036b61076636600461438c565b6001600160a01b031660009081526002602052604090205490565b34801561078d57600080fd5b50610796611681565b60405161035097969594939291906143a9565b3480156107b557600080fd5b5061036b6107c4366004614419565b6116c7565b3480156107d557600080fd5b506107de611799565b60405165ffffffffffff9091168152602001610350565b34801561080157600080fd5b50606461036b565b34801561081557600080fd5b5061036b61082436600461446a565b611821565b34801561083557600080fd5b5061036b611838565b34801561084a57600080fd5b5061031d61085936600461438c565b611852565b34801561086a57600080fd5b50610344610879366004613b10565b611863565b34801561088a57600080fd5b5061036b610899366004613b10565b60009081526004602052604090206001015465ffffffffffff1690565b3480156108c257600080fd5b50600061036b565b3480156108d657600080fd5b506104296108e53660046144c2565b61186c565b3480156108f657600080fd5b5061036b610905366004613b10565b6118b0565b61031d610918366004614555565b6118f3565b34801561092957600080fd5b5061036b610938366004613e4b565b611973565b34801561094957600080fd5b50600b546001600160a01b03166103f1565b34801561096757600080fd5b5060085461036b565b34801561097c57600080fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e908201526103ae565b3480156109c257600080fd5b5061031d6109d1366004613b10565b6119ad565b3480156109e257600080fd5b5061036b7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b348015610a1657600080fd5b5061036b610a25366004614598565b611b49565b348015610a3657600080fd5b50610429610a453660046145c4565b611b6a565b348015610a5657600080fd5b5061036b610a65366004613b10565b611bae565b348015610a7657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006103f1565b61031d610aab366004613b10565b611bb9565b6000610ac4600b546001600160a01b031690565b905090565b60006001600160e01b031982166332a2ad4360e11b1480610afa57506001600160e01b03198216630271189760e51b145b80610b1557506301ffc9a760e01b6001600160e01b03198316145b92915050565b610b23611d55565b610b2c81611dcf565b50565b606060038054610b3e9061462c565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6a9061462c565b8015610bb75780601f10610b8c57610100808354040283529160200191610bb7565b820191906000526020600020905b815481529060010190602001808311610b9a57829003601f168201915b5050505050905090565b600030610bcc610ab0565b6001600160a01b031614610bf357604051637485328f60e11b815260040160405180910390fd5b50630a85bd0160e11b949350505050565b600080610c1386868686611973565b9050610c2881610c236004611e65565b611e88565b506000610c388288888888611ec7565b905065ffffffffffff811615610cb057600082815260046020908152604091829020600101805465ffffffffffff191665ffffffffffff85169081179091558251858152918201527f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892910160405180910390a1610cc9565b604051634844252360e11b815260040160405180910390fd5b5095945050505050565b60008181526009602090815260408083208151815460a0948102820185019093526080810183815260609586958695919485949390928492849190840182828015610d4757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d29575b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015610d9f57602002820191906000526020600020905b815481526020019060010190808311610d8b575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020016000905b82821015610e79578382906000526020600020018054610dec9061462c565b80601f0160208091040260200160405190810160405280929190818152602001828054610e189061462c565b8015610e655780601f10610e3a57610100808354040283529160200191610e65565b820191906000526020600020905b815481529060010190602001808311610e4857829003601f168201915b505050505081526020019060010190610dcd565b505050508152602001600382015481525050905080606001516000801b03610ebc57604051636ad0607560e01b8152600481018790526024015b60405180910390fd5b80516020820151604083015160609093015191989097509195509350915050565b600080610eec86868686611973565b9050610f0c81610efc6005611e65565b610f066004611e65565b17611e88565b506000818152600460205260409020805460ff60f01b1916600160f01b17905530610f35610ab0565b6001600160a01b031614610fbf5760005b8651811015610fbd57306001600160a01b0316878281518110610f6b57610f6b614666565b60200260200101516001600160a01b031603610fb557610fb5858281518110610f9657610f96614666565b6020026020010151805190602001206005611ed690919063ffffffff16565b600101610f46565b505b610fcc8187878787611f38565b30610fd5610ab0565b6001600160a01b03161415801561100157506005546001600160801b03808216600160801b9092041614155b1561100c5760006005555b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b600090815260046020526040902054600160a01b900465ffffffffffff1690565b6000606080606060006008868154811061108657611086614666565b9060005260206000200154945061109c85610cd3565b979992985090969095509350915050565b6000610b1582611f4c565b60008181526009602090815260409182902080548351818402810184019094528084529092611254929091849183018282801561111e57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611100575b50505050508260010180548060200260200160405190810160405280929190818152602001828054801561117157602002820191906000526020600020905b81548152602001906001019080831161115d575b505050505083600201805480602002602001604051908101604052809291908181526020016000905b828210156112465783829060005260206000200180546111b99061462c565b80601f01602080910402602001604051908101604052809291908181526020018280546111e59061462c565b80156112325780601f1061120757610100808354040283529160200191611232565b820191906000526020600020905b81548152906001019060200180831161121557829003601f168201915b50505050508152602001906001019061119a565b505050508460030154611259565b505050565b60008061126886868686611973565b905061127881610c236000611e65565b506000818152600460205260409020546001600160a01b031633146112b25760405163233d98e360e01b8152336004820152602401610eb3565b6112be8686868661208b565b9695505050505050565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561134957506040513d6000823e601f3d908101601f19168201604052611346919081019061467c565b60015b611385575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b6000803390506113ab848285604051806020016040528060008152506120a2565b949350505050565b600080611496876114907f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c6114088e6001600160a01b0316600090815260026020526040902080546001810190915590565b8d8d6040516114189291906146e9565b60405180910390208c805190602001206040516020016114759796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b604051602081830303815290604052805190602001206120c5565b856120f2565b9050806114c1576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610eb3565b61150589888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b9250612166915050565b9998505050505050505050565b60008033905061155c87828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250612166915050565b979650505050505050565b600a8054600091829061157b60018461470f565b8154811061158b5761158b614666565b6000918252602090912001805490915065ffffffffffff811690600160301b90046001600160d01b03168582116115ce576001600160d01b031695945050505050565b6115e26115da87612248565b600a9061227f565b6001600160d01b03169695505050505050565b6000803390506112be86828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506120a292505050565b60003361164a8184612334565b6116725760405163d9b3955760e01b81526001600160a01b0382166004820152602401610eb3565b600061155c87878787866123bb565b6000606080600080600060606116956123ca565b61169d6123f7565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b600080611753846114907ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d789898961171c8b6001600160a01b0316600090815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001611475565b90508061177e576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610eb3565b6112be868587604051806020016040528060008152506120a2565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611815575060408051601f3d908101601f1916820190925261181291810190614729565b60015b61138557610ac4612424565b600061182e84848461242f565b90505b9392505050565b6000611844600a6124c5565b6001600160d01b0316905090565b61185a611d55565b610b2c816124fe565b60006001610b15565b600030611877610ab0565b6001600160a01b03161461189e57604051637485328f60e11b815260040160405180910390fd5b5063bc197c8160e01b95945050505050565b6000818152600460205260408120546118e590600160d01b810463ffffffff1690600160a01b900465ffffffffffff16614751565b65ffffffffffff1692915050565b6118fb611d55565b600080856001600160a01b03168585856040516119199291906146e9565b60006040518083038185875af1925050503d8060008114611956576040519150601f19603f3d011682016040523d82523d6000602084013e61195b565b606091505b509150915061196a8282612567565b50505050505050565b60008484848460405160200161198c9493929190613faa565b60408051601f19818403018152919052805160209091012095945050505050565b600081815260096020908152604091829020805483518184028101840190945280845290926112549290918491830182828015611a1357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116119f5575b505050505082600101805480602002602001604051908101604052809291908181526020018280548015611a6657602002820191906000526020600020905b815481526020019060010190808311611a52575b505050505083600201805480602002602001604051908101604052809291908181526020016000905b82821015611b3b578382906000526020600020018054611aae9061462c565b80601f0160208091040260200160405190810160405280929190818152602001828054611ada9061462c565b8015611b275780601f10611afc57610100808354040283529160200191611b27565b820191906000526020600020905b815481529060010190602001808311611b0a57829003601f168201915b505050505081526020019060010190611a8f565b505050508460030154610c04565b60006118318383611b6560408051602081019091526000815290565b61242f565b600030611b75610ab0565b6001600160a01b031614611b9c57604051637485328f60e11b815260040160405180910390fd5b5063f23a6e6160e01b95945050505050565b6000610b1582612583565b600081815260096020908152604091829020805483518184028101840190945280845290926112549290918491830182828015611c1f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611c01575b505050505082600101805480602002602001604051908101604052809291908181526020018280548015611c7257602002820191906000526020600020905b815481526020019060010190808311611c5e575b505050505083600201805480602002602001604051908101604052809291908181526020016000905b82821015611d47578382906000526020600020018054611cba9061462c565b80601f0160208091040260200160405190810160405280929190818152602001828054611ce69061462c565b8015611d335780601f10611d0857610100808354040283529160200191611d33565b820191906000526020600020905b815481529060010190602001808311611d1657829003601f168201915b505050505081526020019060010190611c9b565b505050508460030154610edd565b33611d5e610ab0565b6001600160a01b031614611d87576040516347096e4760e01b8152336004820152602401610eb3565b30611d90610ab0565b6001600160a01b031614611dcd5760008036604051611db09291906146e9565b604051809103902090505b80611dc6600561262d565b03611dbb57505b565b606480821115611dfc5760405163243e544560e01b81526004810183905260248101829052604401610eb3565b6000611e06611838565b9050611e25611e13611799565b611e1c8561269c565b600a91906126d0565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b6000816007811115611e7957611e79614047565b600160ff919091161b92915050565b600080611e94846110ad565b9050600083611ea283611e65565b1603611831578381846040516331b75e4d60e01b8152600401610eb393929190614770565b60006112be86868686866126eb565b81546001600160801b03600160801b820481169181166001830190911603611f0257611f026041612896565b6001600160801b03808216600090815260018086016020526040909120939093558354919092018216600160801b029116179055565b611f4585858585856128a8565b5050505050565b600080611f588361293e565b90506005816007811115611f6e57611f6e614047565b14611f795792915050565b6000838152600c60205260409081902054600b549151632c258a9f60e11b81526004810182905290916001600160a01b03169063584b153e90602401602060405180830381865afa158015611fd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ff69190614792565b15612005575060059392505050565b600b54604051632ab0f52960e01b8152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa15801561204e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120729190614792565b15612081575060079392505050565b5060029392505050565b600061209985858585612a78565b95945050505050565b6000612099858585856120c060408051602081019091526000815290565b612166565b6000610b156120d2612b17565b8360405161190160f01b8152600281019290925260228201526042902090565b6000836001600160a01b03163b600003612154576000806121138585612c42565b509092509050600081600381111561212d5761212d614047565b14801561214b5750856001600160a01b0316826001600160a01b0316145b92505050611831565b61215f848484612c8f565b9050611831565b600061217686610c236001611e65565b50600061218c8661218689611049565b8561242f565b9050600061219d8888888588612d6a565b905083516000036121f457866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4898884896040516121e794939291906147b4565b60405180910390a261155c565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516122359594939291906147dc565b60405180910390a2979650505050505050565b600065ffffffffffff82111561227b576040516306dfcc6560e41b81526030600482015260248101839052604401610eb3565b5090565b8154600090818160058111156122de57600061229a84612e6d565b6122a4908561470f565b60008881526020902090915081015465ffffffffffff90811690871610156122ce578091506122dc565b6122d9816001614816565b92505b505b60006122ec87878585612fc6565b90508015612327576123118761230360018461470f565b600091825260209091200190565b54600160301b90046001600160d01b031661155c565b6000979650505050505050565b8051600090603481101561234c576001915050610b15565b60131981840101516001600160b01b03198116692370726f706f7365723d60b01b1461237d57600192505050610b15565b60008061238e86602a860386613028565b9150915081158061155c5750866001600160a01b0316816001600160a01b03161494505050505092915050565b60006112be86868686866130d7565b6060610ac47f000000000000000000000000000000000000000000000000000000000000000060006131b2565b6060610ac47f000000000000000000000000000000000000000000000000000000000000000060016131b2565b6000610ac443612248565b60007f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa1580156124a1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061182e9190614829565b805460009080156124f5576124df8361230360018461470f565b54600160301b90046001600160d01b0316611831565b60009392505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b60608261257c576125778261325d565b610b15565b5080610b15565b6000606461259083611567565b604051632394e7a360e21b8152600481018590526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa1580156125f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126199190614829565b6126239190614842565b610b15919061486f565b80546000906001600160801b0380821691600160801b9004168103612656576126566031612896565b6001600160801b038181166000908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b60006001600160d01b0382111561227b576040516306dfcc6560e41b815260d0600482015260248101839052604401610eb3565b6000806126de858585613286565b915091505b935093915050565b600080600b60009054906101000a90046001600160a01b03166001600160a01b031663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa158015612741573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127659190614829565b905060003060601b6bffffffffffffffffffffffff19168418600b5460405163b1c5f42760e01b81529192506001600160a01b03169063b1c5f427906127b8908a908a908a906000908890600401614891565b602060405180830381865afa1580156127d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f99190614829565b6000898152600c602052604080822092909255600b5491516308f2a0bb60e41b81526001600160a01b0390921691638f2a0bb091612844918b918b918b919088908a906004016148df565b600060405180830381600087803b15801561285e57600080fd5b505af1158015612872573d6000803e3d6000fd5b5050505061288a82426128859190614816565b612248565b98975050505050505050565b634e487b71600052806020526024601cfd5b600b546001600160a01b031663e38335e53486868660003060601b6bffffffffffffffffffffffff191688186040518763ffffffff1660e01b81526004016128f4959493929190614891565b6000604051808303818588803b15801561290d57600080fd5b505af1158015612921573d6000803e3d6000fd5b50505060009687525050600c602052505060408320929092555050565b6000818152600460205260408120805460ff600160f01b8204811691600160f81b900416811561297357506007949350505050565b801561298457506002949350505050565b600061298f86611049565b9050806000036129b557604051636ad0607560e01b815260048101879052602401610eb3565b60006129bf611799565b65ffffffffffff1690508082106129dd575060009695505050505050565b60006129e8886118b0565b90508181106129ff57506001979650505050505050565b612a08886133da565b1580612a2857506000888152600760205260409020805460019091015411155b15612a3b57506003979650505050505050565b60008881526004602052604090206001015465ffffffffffff16600003612a6a57506004979650505050505050565b506005979650505050505050565b600080612a8786868686613411565b6000818152600c60205260409020549091508015610cc957600b5460405163c4d252f560e01b8152600481018390526001600160a01b039091169063c4d252f590602401600060405180830381600087803b158015612ae557600080fd5b505af1158015612af9573d6000803e3d6000fd5b5050506000838152600c602052604081205550509050949350505050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015612b7057507f000000000000000000000000000000000000000000000000000000000000000046145b15612b9a57507f000000000000000000000000000000000000000000000000000000000000000090565b610ac4604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a082015260009060c00160405160208183030381529060405280519060200120905090565b60008060008351604103612c7c5760208401516040850151606086015160001a612c6e888285856134c2565b955095509550505050612c88565b50508151600091506002905b9250925092565b6000806000856001600160a01b03168585604051602401612cb1929190614937565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b17905251612ce69190614950565b600060405180830381855afa9150503d8060008114612d21576040519150601f19603f3d011682016040523d82523d6000602084013e612d26565b606091505b5091509150818015612d3a57506020815110155b80156112be57508051630b135d3f60e11b90612d5f9083016020908101908401614829565b149695505050505050565b60008581526007602090815260408083206001600160a01b03881684526003810190925282205460ff1615612dbd576040516371c6af4960e01b81526001600160a01b0387166004820152602401610eb3565b6001600160a01b03861660009081526003820160205260409020805460ff1916600117905560ff8516612e095783816000016000828254612dfe9190614816565b90915550612e629050565b60001960ff861601612e295783816001016000828254612dfe9190614816565b60011960ff861601612e495783816002016000828254612dfe9190614816565b6040516303599be160e11b815260040160405180910390fd5b509195945050505050565b600060018211612e7b575090565b816001600160801b8210612e945760809190911c9060401b5b680100000000000000008210612eaf5760409190911c9060201b5b6401000000008210612ec65760209190911c9060101b5b620100008210612edb5760109190911c9060081b5b6101008210612eef5760089190911c9060041b5b60108210612f025760049190911c9060021b5b60048210612f0e5760011b5b600302600190811c90818581612f2657612f26614859565b048201901c90506001818581612f3e57612f3e614859565b048201901c90506001818581612f5657612f56614859565b048201901c90506001818581612f6e57612f6e614859565b048201901c90506001818581612f8657612f86614859565b048201901c90506001818581612f9e57612f9e614859565b048201901c9050612fbd818581612fb757612fb7614859565b04821190565b90039392505050565b60005b81831015613020576000612fdd8484613591565b60008781526020902090915065ffffffffffff86169082015465ffffffffffff16111561300c5780925061301a565b613017816001614816565b93505b50612fc9565b509392505050565b600080845183118061303957508284115b15613049575060009050806126e3565b6000613056856001614816565b8411801561307e575061060f60f31b6130728787016020015190565b6001600160f01b031916145b9050600061308f8215156002614842565b61309a906028614816565b9050806130a7878761470f565b036130c9576000806130ba8989896135ac565b90965094506126e39350505050565b6000809350935050506126e3565b6000806130e78787878787613672565b6008805460018101825560009182527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee301829055604080516080810182528a815260208181018b90528183018a905288518982012060608301528484526009815291909220825180519495509293909261316592849291019061398c565b50602082810151805161317e92600185019201906139ed565b506040820151805161319a916002840191602090910190613a28565b50606091909101516003909101559695505050505050565b606060ff83146131cc576131c58361387b565b9050610b15565b8180546131d89061462c565b80601f01602080910402602001604051908101604052809291908181526020018280546132049061462c565b80156132515780601f1061322657610100808354040283529160200191613251565b820191906000526020600020905b81548152906001019060200180831161323457829003601f168201915b50505050509050610b15565b80511561326d5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b82546000908190801561337f5760006132a48761230360018561470f565b805490915065ffffffffffff80821691600160301b90046001600160d01b03169088168211156132e757604051632520601d60e01b815260040160405180910390fd5b8765ffffffffffff168265ffffffffffff160361332057825465ffffffffffff16600160301b6001600160d01b03891602178355613371565b6040805180820190915265ffffffffffff808a1682526001600160d01b03808a1660208085019182528d54600181018f5560008f81529190912094519151909216600160301b029216919091179101555b94508593506126e392505050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a5560008a815291822095519251909316600160301b0291909316179201919091559050816126e3565b6000818152600760205260408120600281015460018201546133fc9190614816565b613408610a6585611049565b11159392505050565b60008061342086868686611973565b905061346e816134306007611e65565b61343a6006611e65565b6134446002611e65565b600161345160078261496c565b61345c906002614a69565b613466919061470f565b181818611e88565b506000818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906110389083815260200190565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156134fd5750600091506003905082613587565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613551573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661357d57506000925060019150829050613587565b9250600091508190505b9450945094915050565b60006135a0600284841861486f565b61183190848416614816565b60008084816135bc866001614816565b851180156135e4575061060f60f31b6135d88388016020015190565b6001600160f01b031916145b905060006135f58215156002614842565b9050600080613604838a614816565b90505b878110156136615760006136266136218784016020015190565b6138ba565b9050600f8160ff16111561364657600080975097505050505050506126e3565b613651601084614842565b60ff909116019150600101613607565b506001999098509650505050505050565b60006136878686868680519060200120611973565b90508451865114158061369c57508351865114155b806136a657508551155b156136db57855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610eb3565b600081815260046020526040902054600160a01b900465ffffffffffff16156137265780613708826110ad565b6040516331b75e4d60e01b8152610eb3929190600090600401614770565b600062015180613734611799565b65ffffffffffff166137469190614816565b9050600062093a80600084815260046020526040902080546001600160a01b0319166001600160a01b03871617815590915061378183612248565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b199091161781556137ae82613933565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b0381111561381157613811613ba1565b60405190808252806020026020018201604052801561384457816020015b606081526020019060019003908161382f5790505b508c896138518a82614816565b8e60405161386799989796959493929190614a78565b60405180910390a150505095945050505050565b6060600061388883613964565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b600060f882901c602f811180156138d45750603a8160ff16105b156138e257602f1901610b15565b60608160ff161180156138f8575060678160ff16105b156139065760561901610b15565b60408160ff1611801561391c575060478160ff16105b1561392a5760361901610b15565b5060ff92915050565b600063ffffffff82111561227b576040516306dfcc6560e41b81526020600482015260248101839052604401610eb3565b600060ff8216601f811115610b1557604051632cd44ac360e21b815260040160405180910390fd5b8280548282559060005260206000209081019282156139e1579160200282015b828111156139e157825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906139ac565b5061227b929150613a7a565b8280548282559060005260206000209081019282156139e1579160200282015b828111156139e1578251825591602001919060010190613a0d565b828054828255906000526020600020908101928215613a6e579160200282015b82811115613a6e5782518290613a5e9082614b9f565b5091602001919060010190613a48565b5061227b929150613a8f565b5b8082111561227b5760008155600101613a7b565b8082111561227b576000613aa38282613aac565b50600101613a8f565b508054613ab89061462c565b6000825580601f10613ac8575050565b601f016020900490600052602060002090810190610b2c9190613a7a565b600060208284031215613af857600080fd5b81356001600160e01b03198116811461183157600080fd5b600060208284031215613b2257600080fd5b5035919050565b60005b83811015613b44578181015183820152602001613b2c565b50506000910152565b60008151808452613b65816020860160208601613b29565b601f01601f19169290920160200192915050565b6020815260006118316020830184613b4d565b6001600160a01b0381168114610b2c57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715613bdf57613bdf613ba1565b604052919050565b60006001600160401b03821115613c0057613c00613ba1565b50601f01601f191660200190565b6000613c21613c1c84613be7565b613bb7565b9050828152838383011115613c3557600080fd5b828260208301376000602084830101529392505050565b600082601f830112613c5d57600080fd5b61183183833560208501613c0e565b60008060008060808587031215613c8257600080fd5b8435613c8d81613b8c565b93506020850135613c9d81613b8c565b92506040850135915060608501356001600160401b03811115613cbf57600080fd5b613ccb87828801613c4c565b91505092959194509250565b60006001600160401b03821115613cf057613cf0613ba1565b5060051b60200190565b600082601f830112613d0b57600080fd5b81356020613d1b613c1c83613cd7565b8083825260208201915060208460051b870101935086841115613d3d57600080fd5b602086015b84811015613d62578035613d5581613b8c565b8352918301918301613d42565b509695505050505050565b600082601f830112613d7e57600080fd5b81356020613d8e613c1c83613cd7565b8083825260208201915060208460051b870101935086841115613db057600080fd5b602086015b84811015613d625780358352918301918301613db5565b600082601f830112613ddd57600080fd5b81356020613ded613c1c83613cd7565b82815260059290921b84018101918181019086841115613e0c57600080fd5b8286015b84811015613d625780356001600160401b03811115613e2f5760008081fd5b613e3d8986838b0101613c4c565b845250918301918301613e10565b60008060008060808587031215613e6157600080fd5b84356001600160401b0380821115613e7857600080fd5b613e8488838901613cfa565b95506020870135915080821115613e9a57600080fd5b613ea688838901613d6d565b94506040870135915080821115613ebc57600080fd5b50613ec987828801613dcc565b949793965093946060013593505050565b60008151808452602080850194506020840160005b83811015613f145781516001600160a01b031687529582019590820190600101613eef565b509495945050505050565b60008151808452602080850194506020840160005b83811015613f1457815187529582019590820190600101613f34565b60008282518085526020808601955060208260051b8401016020860160005b84811015613f9d57601f19868403018952613f8b838351613b4d565b98840198925090830190600101613f6f565b5090979650505050505050565b608081526000613fbd6080830187613eda565b8281036020840152613fcf8187613f1f565b90508281036040840152613fe38186613f50565b91505082606083015295945050505050565b85815260a06020820152600061400e60a0830187613eda565b82810360408401526140208187613f1f565b905082810360608401526140348186613f50565b9150508260808301529695505050505050565b634e487b7160e01b600052602160045260246000fd5b6008811061407b57634e487b7160e01b600052602160045260246000fd5b9052565b60208101610b15828461405d565b600080604083850312156140a057600080fd5b8235915060208301356140b281613b8c565b809150509250929050565b803560ff8116811461138557600080fd5b600080604083850312156140e157600080fd5b823591506140f1602084016140bd565b90509250929050565b60008083601f84011261410c57600080fd5b5081356001600160401b0381111561412357600080fd5b60208301915083602082850101111561413b57600080fd5b9250929050565b600080600080600080600060c0888a03121561415d57600080fd5b8735965061416d602089016140bd565b9550604088013561417d81613b8c565b945060608801356001600160401b038082111561419957600080fd5b6141a58b838c016140fa565b909650945060808a01359150808211156141be57600080fd5b6141ca8b838c01613c4c565b935060a08a01359150808211156141e057600080fd5b506141ed8a828b01613c4c565b91505092959891949750929550565b60008060008060006080868803121561421457600080fd5b85359450614224602087016140bd565b935060408601356001600160401b038082111561424057600080fd5b61424c89838a016140fa565b9095509350606088013591508082111561426557600080fd5b5061427288828901613c4c565b9150509295509295909350565b6000806000806060858703121561429557600080fd5b843593506142a5602086016140bd565b925060408501356001600160401b038111156142c057600080fd5b6142cc878288016140fa565b95989497509550505050565b600080600080608085870312156142ee57600080fd5b84356001600160401b038082111561430557600080fd5b61431188838901613cfa565b9550602087013591508082111561432757600080fd5b61433388838901613d6d565b9450604087013591508082111561434957600080fd5b61435588838901613dcc565b9350606087013591508082111561436b57600080fd5b508501601f8101871361437d57600080fd5b613ccb87823560208401613c0e565b60006020828403121561439e57600080fd5b813561183181613b8c565b60ff60f81b8816815260e0602082015260006143c860e0830189613b4d565b82810360408401526143da8189613b4d565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152905061440b8185613f1f565b9a9950505050505050505050565b6000806000806080858703121561442f57600080fd5b8435935061443f602086016140bd565b9250604085013561444f81613b8c565b915060608501356001600160401b03811115613cbf57600080fd5b60008060006060848603121561447f57600080fd5b833561448a81613b8c565b92506020840135915060408401356001600160401b038111156144ac57600080fd5b6144b886828701613c4c565b9150509250925092565b600080600080600060a086880312156144da57600080fd5b85356144e581613b8c565b945060208601356144f581613b8c565b935060408601356001600160401b038082111561451157600080fd5b61451d89838a01613d6d565b9450606088013591508082111561453357600080fd5b61453f89838a01613d6d565b9350608088013591508082111561426557600080fd5b6000806000806060858703121561456b57600080fd5b843561457681613b8c565b93506020850135925060408501356001600160401b038111156142c057600080fd5b600080604083850312156145ab57600080fd5b82356145b681613b8c565b946020939093013593505050565b600080600080600060a086880312156145dc57600080fd5b85356145e781613b8c565b945060208601356145f781613b8c565b9350604086013592506060860135915060808601356001600160401b0381111561462057600080fd5b61427288828901613c4c565b600181811c9082168061464057607f821691505b60208210810361466057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561468e57600080fd5b81516001600160401b038111156146a457600080fd5b8201601f810184136146b557600080fd5b80516146c3613c1c82613be7565b8181528560208385010111156146d857600080fd5b612099826020830160208601613b29565b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610b1557610b156146f9565b5092915050565b60006020828403121561473b57600080fd5b815165ffffffffffff8116811461183157600080fd5b65ffffffffffff818116838216019080821115614722576147226146f9565b83815260608101614784602083018561405d565b826040830152949350505050565b6000602082840312156147a457600080fd5b8151801515811461183157600080fd5b84815260ff841660208201528260408201526080606082015260006112be6080830184613b4d565b85815260ff8516602082015283604082015260a06060820152600061480460a0830185613b4d565b828103608084015261288a8185613b4d565b80820180821115610b1557610b156146f9565b60006020828403121561483b57600080fd5b5051919050565b8082028115828204841417610b1557610b156146f9565b634e487b7160e01b600052601260045260246000fd5b60008261488c57634e487b7160e01b600052601260045260246000fd5b500490565b60a0815260006148a460a0830188613eda565b82810360208401526148b68188613f1f565b905082810360408401526148ca8187613f50565b60608401959095525050608001529392505050565b60c0815260006148f260c0830189613eda565b82810360208401526149048189613f1f565b905082810360408401526149188188613f50565b60608401969096525050608081019290925260a0909101529392505050565b82815260406020820152600061182e6040830184613b4d565b60008251614962818460208701613b29565b9190910192915050565b60ff8181168382160190811115610b1557610b156146f9565b600181815b808511156149c05781600019048211156149a6576149a66146f9565b808516156149b357918102915b93841c939080029061498a565b509250929050565b6000826149d757506001610b15565b816149e457506000610b15565b81600181146149fa5760028114614a0457614a20565b6001915050610b15565b60ff841115614a1557614a156146f9565b50506001821b610b15565b5060208310610133831016604e8410600b8410161715614a43575081810a610b15565b614a4d8383614985565b8060001904821115614a6157614a616146f9565b029392505050565b600061183160ff8416836149c8565b60006101208b8352602060018060a01b038c1681850152816040850152614aa18285018c613eda565b91508382036060850152614ab5828b613f1f565b915083820360808501528189518084528284019150828160051b850101838c0160005b83811015614b0657601f19878403018552614af4838351613b4d565b94860194925090850190600101614ad8565b505086810360a0880152614b1a818c613f50565b9450505050508560c08401528460e0840152828103610100840152614b3f8185613b4d565b9c9b505050505050505050505050565b601f821115611254576000816000526020600020601f850160051c81016020861015614b785750805b601f850160051c820191505b81811015614b9757828155600101614b84565b505050505050565b81516001600160401b03811115614bb857614bb8613ba1565b614bcc81614bc6845461462c565b84614b4f565b602080601f831160018114614c015760008415614be95750858301515b600019600386901b1c1916600185901b178555614b97565b600085815260208120601f198616915b82811015614c3057888601518255948401946001909101908401614c11565b5085821015614c4e5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212209f32ac32ca2c4a01afcb7dae5da054d27ddbe3ce7b896157437bba0aa60d5cd464736f6c63430008160033",
}

// LeeaGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use LeeaGovernanceMetaData.ABI instead.
var LeeaGovernanceABI = LeeaGovernanceMetaData.ABI

// LeeaGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LeeaGovernanceMetaData.Bin instead.
var LeeaGovernanceBin = LeeaGovernanceMetaData.Bin

// DeployLeeaGovernance deploys a new Ethereum contract, binding an instance of LeeaGovernance to it.
func DeployLeeaGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _timelock common.Address) (common.Address, *types.Transaction, *LeeaGovernance, error) {
	parsed, err := LeeaGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LeeaGovernanceBin), backend, _token, _timelock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LeeaGovernance{LeeaGovernanceCaller: LeeaGovernanceCaller{contract: contract}, LeeaGovernanceTransactor: LeeaGovernanceTransactor{contract: contract}, LeeaGovernanceFilterer: LeeaGovernanceFilterer{contract: contract}}, nil
}

// LeeaGovernance is an auto generated Go binding around an Ethereum contract.
type LeeaGovernance struct {
	LeeaGovernanceCaller     // Read-only binding to the contract
	LeeaGovernanceTransactor // Write-only binding to the contract
	LeeaGovernanceFilterer   // Log filterer for contract events
}

// LeeaGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeeaGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeeaGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeeaGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeeaGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeeaGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeeaGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeeaGovernanceSession struct {
	Contract     *LeeaGovernance   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeeaGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeeaGovernanceCallerSession struct {
	Contract *LeeaGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LeeaGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeeaGovernanceTransactorSession struct {
	Contract     *LeeaGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LeeaGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeeaGovernanceRaw struct {
	Contract *LeeaGovernance // Generic contract binding to access the raw methods on
}

// LeeaGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeeaGovernanceCallerRaw struct {
	Contract *LeeaGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// LeeaGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeeaGovernanceTransactorRaw struct {
	Contract *LeeaGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeeaGovernance creates a new instance of LeeaGovernance, bound to a specific deployed contract.
func NewLeeaGovernance(address common.Address, backend bind.ContractBackend) (*LeeaGovernance, error) {
	contract, err := bindLeeaGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernance{LeeaGovernanceCaller: LeeaGovernanceCaller{contract: contract}, LeeaGovernanceTransactor: LeeaGovernanceTransactor{contract: contract}, LeeaGovernanceFilterer: LeeaGovernanceFilterer{contract: contract}}, nil
}

// NewLeeaGovernanceCaller creates a new read-only instance of LeeaGovernance, bound to a specific deployed contract.
func NewLeeaGovernanceCaller(address common.Address, caller bind.ContractCaller) (*LeeaGovernanceCaller, error) {
	contract, err := bindLeeaGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceCaller{contract: contract}, nil
}

// NewLeeaGovernanceTransactor creates a new write-only instance of LeeaGovernance, bound to a specific deployed contract.
func NewLeeaGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*LeeaGovernanceTransactor, error) {
	contract, err := bindLeeaGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceTransactor{contract: contract}, nil
}

// NewLeeaGovernanceFilterer creates a new log filterer instance of LeeaGovernance, bound to a specific deployed contract.
func NewLeeaGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*LeeaGovernanceFilterer, error) {
	contract, err := bindLeeaGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceFilterer{contract: contract}, nil
}

// bindLeeaGovernance binds a generic wrapper to an already deployed contract.
func bindLeeaGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LeeaGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeeaGovernance *LeeaGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeeaGovernance.Contract.LeeaGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeeaGovernance *LeeaGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.LeeaGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeeaGovernance *LeeaGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.LeeaGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeeaGovernance *LeeaGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeeaGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeeaGovernance *LeeaGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeeaGovernance *LeeaGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _LeeaGovernance.Contract.BALLOTTYPEHASH(&_LeeaGovernance.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _LeeaGovernance.Contract.BALLOTTYPEHASH(&_LeeaGovernance.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_LeeaGovernance *LeeaGovernanceSession) CLOCKMODE() (string, error) {
	return _LeeaGovernance.Contract.CLOCKMODE(&_LeeaGovernance.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCallerSession) CLOCKMODE() (string, error) {
	return _LeeaGovernance.Contract.CLOCKMODE(&_LeeaGovernance.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_LeeaGovernance *LeeaGovernanceCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_LeeaGovernance *LeeaGovernanceSession) COUNTINGMODE() (string, error) {
	return _LeeaGovernance.Contract.COUNTINGMODE(&_LeeaGovernance.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_LeeaGovernance *LeeaGovernanceCallerSession) COUNTINGMODE() (string, error) {
	return _LeeaGovernance.Contract.COUNTINGMODE(&_LeeaGovernance.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _LeeaGovernance.Contract.EXTENDEDBALLOTTYPEHASH(&_LeeaGovernance.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_LeeaGovernance *LeeaGovernanceCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _LeeaGovernance.Contract.EXTENDEDBALLOTTYPEHASH(&_LeeaGovernance.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_LeeaGovernance *LeeaGovernanceCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_LeeaGovernance *LeeaGovernanceSession) Clock() (*big.Int, error) {
	return _LeeaGovernance.Contract.Clock(&_LeeaGovernance.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Clock() (*big.Int, error) {
	return _LeeaGovernance.Contract.Clock(&_LeeaGovernance.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_LeeaGovernance *LeeaGovernanceCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_LeeaGovernance *LeeaGovernanceSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _LeeaGovernance.Contract.Eip712Domain(&_LeeaGovernance.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _LeeaGovernance.Contract.Eip712Domain(&_LeeaGovernance.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.GetVotes(&_LeeaGovernance.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.GetVotes(&_LeeaGovernance.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _LeeaGovernance.Contract.GetVotesWithParams(&_LeeaGovernance.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _LeeaGovernance.Contract.GetVotesWithParams(&_LeeaGovernance.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _LeeaGovernance.Contract.HasVoted(&_LeeaGovernance.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _LeeaGovernance.Contract.HasVoted(&_LeeaGovernance.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _LeeaGovernance.Contract.HashProposal(&_LeeaGovernance.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _LeeaGovernance.Contract.HashProposal(&_LeeaGovernance.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LeeaGovernance *LeeaGovernanceSession) Name() (string, error) {
	return _LeeaGovernance.Contract.Name(&_LeeaGovernance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Name() (string, error) {
	return _LeeaGovernance.Contract.Name(&_LeeaGovernance.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LeeaGovernance.Contract.Nonces(&_LeeaGovernance.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LeeaGovernance.Contract.Nonces(&_LeeaGovernance.CallOpts, owner)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalCount() (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalCount(&_LeeaGovernance.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalCount() (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalCount(&_LeeaGovernance.CallOpts)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalDeadline(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalDeadline(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalDetails is a free data retrieval call binding the contract method 0x16e9eaec.
//
// Solidity: function proposalDetails(uint256 proposalId) view returns(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalDetails(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalDetails", proposalId)

	outstruct := new(struct {
		Targets         []common.Address
		Values          []*big.Int
		Calldatas       [][]byte
		DescriptionHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Targets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Calldatas = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)
	outstruct.DescriptionHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ProposalDetails is a free data retrieval call binding the contract method 0x16e9eaec.
//
// Solidity: function proposalDetails(uint256 proposalId) view returns(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalDetails(proposalId *big.Int) (struct {
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	return _LeeaGovernance.Contract.ProposalDetails(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalDetails is a free data retrieval call binding the contract method 0x16e9eaec.
//
// Solidity: function proposalDetails(uint256 proposalId) view returns(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalDetails(proposalId *big.Int) (struct {
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	return _LeeaGovernance.Contract.ProposalDetails(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalDetailsAt is a free data retrieval call binding the contract method 0x2e82db94.
//
// Solidity: function proposalDetailsAt(uint256 index) view returns(uint256 proposalId, address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalDetailsAt(opts *bind.CallOpts, index *big.Int) (struct {
	ProposalId      *big.Int
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalDetailsAt", index)

	outstruct := new(struct {
		ProposalId      *big.Int
		Targets         []common.Address
		Values          []*big.Int
		Calldatas       [][]byte
		DescriptionHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Targets = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Calldatas = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)
	outstruct.DescriptionHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// ProposalDetailsAt is a free data retrieval call binding the contract method 0x2e82db94.
//
// Solidity: function proposalDetailsAt(uint256 index) view returns(uint256 proposalId, address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalDetailsAt(index *big.Int) (struct {
	ProposalId      *big.Int
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	return _LeeaGovernance.Contract.ProposalDetailsAt(&_LeeaGovernance.CallOpts, index)
}

// ProposalDetailsAt is a free data retrieval call binding the contract method 0x2e82db94.
//
// Solidity: function proposalDetailsAt(uint256 index) view returns(uint256 proposalId, address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalDetailsAt(index *big.Int) (struct {
	ProposalId      *big.Int
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	DescriptionHash [32]byte
}, error) {
	return _LeeaGovernance.Contract.ProposalDetailsAt(&_LeeaGovernance.CallOpts, index)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalEta(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalEta(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalNeedsQueuing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalNeedsQueuing", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _LeeaGovernance.Contract.ProposalNeedsQueuing(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _LeeaGovernance.Contract.ProposalNeedsQueuing(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _LeeaGovernance.Contract.ProposalProposer(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _LeeaGovernance.Contract.ProposalProposer(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalSnapshot(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalSnapshot(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalThreshold() (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalThreshold(&_LeeaGovernance.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalThreshold() (*big.Int, error) {
	return _LeeaGovernance.Contract.ProposalThreshold(&_LeeaGovernance.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_LeeaGovernance *LeeaGovernanceCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_LeeaGovernance *LeeaGovernanceSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _LeeaGovernance.Contract.ProposalVotes(&_LeeaGovernance.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_LeeaGovernance *LeeaGovernanceCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _LeeaGovernance.Contract.ProposalVotes(&_LeeaGovernance.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.Quorum(&_LeeaGovernance.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.Quorum(&_LeeaGovernance.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) QuorumDenominator() (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumDenominator(&_LeeaGovernance.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) QuorumDenominator() (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumDenominator(&_LeeaGovernance.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumNumerator(&_LeeaGovernance.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumNumerator(&_LeeaGovernance.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) QuorumNumerator0() (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumNumerator0(&_LeeaGovernance.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _LeeaGovernance.Contract.QuorumNumerator0(&_LeeaGovernance.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_LeeaGovernance *LeeaGovernanceCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_LeeaGovernance *LeeaGovernanceSession) State(proposalId *big.Int) (uint8, error) {
	return _LeeaGovernance.Contract.State(&_LeeaGovernance.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_LeeaGovernance *LeeaGovernanceCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _LeeaGovernance.Contract.State(&_LeeaGovernance.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LeeaGovernance.Contract.SupportsInterface(&_LeeaGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LeeaGovernance *LeeaGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LeeaGovernance.Contract.SupportsInterface(&_LeeaGovernance.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_LeeaGovernance *LeeaGovernanceCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_LeeaGovernance *LeeaGovernanceSession) Timelock() (common.Address, error) {
	return _LeeaGovernance.Contract.Timelock(&_LeeaGovernance.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Timelock() (common.Address, error) {
	return _LeeaGovernance.Contract.Timelock(&_LeeaGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LeeaGovernance *LeeaGovernanceCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LeeaGovernance *LeeaGovernanceSession) Token() (common.Address, error) {
	return _LeeaGovernance.Contract.Token(&_LeeaGovernance.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Token() (common.Address, error) {
	return _LeeaGovernance.Contract.Token(&_LeeaGovernance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LeeaGovernance *LeeaGovernanceSession) Version() (string, error) {
	return _LeeaGovernance.Contract.Version(&_LeeaGovernance.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LeeaGovernance *LeeaGovernanceCallerSession) Version() (string, error) {
	return _LeeaGovernance.Contract.Version(&_LeeaGovernance.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) VotingDelay() (*big.Int, error) {
	return _LeeaGovernance.Contract.VotingDelay(&_LeeaGovernance.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) VotingDelay() (*big.Int, error) {
	return _LeeaGovernance.Contract.VotingDelay(&_LeeaGovernance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LeeaGovernance.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) VotingPeriod() (*big.Int, error) {
	return _LeeaGovernance.Contract.VotingPeriod(&_LeeaGovernance.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() pure returns(uint256)
func (_LeeaGovernance *LeeaGovernanceCallerSession) VotingPeriod() (*big.Int, error) {
	return _LeeaGovernance.Contract.VotingPeriod(&_LeeaGovernance.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) Cancel(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "cancel", proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Cancel(&_LeeaGovernance.TransactOpts, proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Cancel(&_LeeaGovernance.TransactOpts, proposalId)
}

// Cancel0 is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) Cancel0(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "cancel0", targets, values, calldatas, descriptionHash)
}

// Cancel0 is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Cancel0(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Cancel0(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel0 is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Cancel0(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Cancel0(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVote(&_LeeaGovernance.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVote(&_LeeaGovernance.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteBySig(&_LeeaGovernance.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteBySig(&_LeeaGovernance.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReason(&_LeeaGovernance.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReason(&_LeeaGovernance.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReasonAndParams(&_LeeaGovernance.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReasonAndParams(&_LeeaGovernance.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReasonAndParamsBySig(&_LeeaGovernance.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.CastVoteWithReasonAndParamsBySig(&_LeeaGovernance.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Execute(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Execute(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) Execute0(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "execute0", proposalId)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_LeeaGovernance *LeeaGovernanceSession) Execute0(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Execute0(&_LeeaGovernance.TransactOpts, proposalId)
}

// Execute0 is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Execute0(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Execute0(&_LeeaGovernance.TransactOpts, proposalId)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC1155BatchReceived(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC1155BatchReceived(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC1155Received(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC1155Received(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC721Received(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.OnERC721Received(&_LeeaGovernance.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Propose(&_LeeaGovernance.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Propose(&_LeeaGovernance.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Queue(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Queue(&_LeeaGovernance.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) Queue0(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "queue0", proposalId)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceSession) Queue0(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Queue0(&_LeeaGovernance.TransactOpts, proposalId)
}

// Queue0 is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Queue0(proposalId *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Queue0(&_LeeaGovernance.TransactOpts, proposalId)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_LeeaGovernance *LeeaGovernanceSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Relay(&_LeeaGovernance.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Relay(&_LeeaGovernance.TransactOpts, target, value, data)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_LeeaGovernance *LeeaGovernanceSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.UpdateQuorumNumerator(&_LeeaGovernance.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.UpdateQuorumNumerator(&_LeeaGovernance.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _LeeaGovernance.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_LeeaGovernance *LeeaGovernanceSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.UpdateTimelock(&_LeeaGovernance.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _LeeaGovernance.Contract.UpdateTimelock(&_LeeaGovernance.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeeaGovernance.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeeaGovernance *LeeaGovernanceSession) Receive() (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Receive(&_LeeaGovernance.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LeeaGovernance *LeeaGovernanceTransactorSession) Receive() (*types.Transaction, error) {
	return _LeeaGovernance.Contract.Receive(&_LeeaGovernance.TransactOpts)
}

// LeeaGovernanceEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the LeeaGovernance contract.
type LeeaGovernanceEIP712DomainChangedIterator struct {
	Event *LeeaGovernanceEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceEIP712DomainChanged represents a EIP712DomainChanged event raised by the LeeaGovernance contract.
type LeeaGovernanceEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*LeeaGovernanceEIP712DomainChangedIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceEIP712DomainChangedIterator{contract: _LeeaGovernance.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceEIP712DomainChanged)
				if err := _LeeaGovernance.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseEIP712DomainChanged(log types.Log) (*LeeaGovernanceEIP712DomainChanged, error) {
	event := new(LeeaGovernanceEIP712DomainChanged)
	if err := _LeeaGovernance.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the LeeaGovernance contract.
type LeeaGovernanceProposalCanceledIterator struct {
	Event *LeeaGovernanceProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceProposalCanceled represents a ProposalCanceled event raised by the LeeaGovernance contract.
type LeeaGovernanceProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*LeeaGovernanceProposalCanceledIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceProposalCanceledIterator{contract: _LeeaGovernance.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceProposalCanceled)
				if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseProposalCanceled(log types.Log) (*LeeaGovernanceProposalCanceled, error) {
	event := new(LeeaGovernanceProposalCanceled)
	if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the LeeaGovernance contract.
type LeeaGovernanceProposalCreatedIterator struct {
	Event *LeeaGovernanceProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceProposalCreated represents a ProposalCreated event raised by the LeeaGovernance contract.
type LeeaGovernanceProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*LeeaGovernanceProposalCreatedIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceProposalCreatedIterator{contract: _LeeaGovernance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceProposalCreated) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceProposalCreated)
				if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseProposalCreated(log types.Log) (*LeeaGovernanceProposalCreated, error) {
	event := new(LeeaGovernanceProposalCreated)
	if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the LeeaGovernance contract.
type LeeaGovernanceProposalExecutedIterator struct {
	Event *LeeaGovernanceProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceProposalExecuted represents a ProposalExecuted event raised by the LeeaGovernance contract.
type LeeaGovernanceProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*LeeaGovernanceProposalExecutedIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceProposalExecutedIterator{contract: _LeeaGovernance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceProposalExecuted)
				if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseProposalExecuted(log types.Log) (*LeeaGovernanceProposalExecuted, error) {
	event := new(LeeaGovernanceProposalExecuted)
	if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the LeeaGovernance contract.
type LeeaGovernanceProposalQueuedIterator struct {
	Event *LeeaGovernanceProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceProposalQueued represents a ProposalQueued event raised by the LeeaGovernance contract.
type LeeaGovernanceProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*LeeaGovernanceProposalQueuedIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceProposalQueuedIterator{contract: _LeeaGovernance.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceProposalQueued) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceProposalQueued)
				if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseProposalQueued(log types.Log) (*LeeaGovernanceProposalQueued, error) {
	event := new(LeeaGovernanceProposalQueued)
	if err := _LeeaGovernance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the LeeaGovernance contract.
type LeeaGovernanceQuorumNumeratorUpdatedIterator struct {
	Event *LeeaGovernanceQuorumNumeratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceQuorumNumeratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceQuorumNumeratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the LeeaGovernance contract.
type LeeaGovernanceQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*LeeaGovernanceQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceQuorumNumeratorUpdatedIterator{contract: _LeeaGovernance.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceQuorumNumeratorUpdated)
				if err := _LeeaGovernance.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*LeeaGovernanceQuorumNumeratorUpdated, error) {
	event := new(LeeaGovernanceQuorumNumeratorUpdated)
	if err := _LeeaGovernance.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the LeeaGovernance contract.
type LeeaGovernanceTimelockChangeIterator struct {
	Event *LeeaGovernanceTimelockChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceTimelockChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceTimelockChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceTimelockChange represents a TimelockChange event raised by the LeeaGovernance contract.
type LeeaGovernanceTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*LeeaGovernanceTimelockChangeIterator, error) {

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceTimelockChangeIterator{contract: _LeeaGovernance.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceTimelockChange) (event.Subscription, error) {

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceTimelockChange)
				if err := _LeeaGovernance.contract.UnpackLog(event, "TimelockChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockChange is a log parse operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseTimelockChange(log types.Log) (*LeeaGovernanceTimelockChange, error) {
	event := new(LeeaGovernanceTimelockChange)
	if err := _LeeaGovernance.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the LeeaGovernance contract.
type LeeaGovernanceVoteCastIterator struct {
	Event *LeeaGovernanceVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceVoteCast represents a VoteCast event raised by the LeeaGovernance contract.
type LeeaGovernanceVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*LeeaGovernanceVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceVoteCastIterator{contract: _LeeaGovernance.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceVoteCast)
				if err := _LeeaGovernance.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseVoteCast(log types.Log) (*LeeaGovernanceVoteCast, error) {
	event := new(LeeaGovernanceVoteCast)
	if err := _LeeaGovernance.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeeaGovernanceVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the LeeaGovernance contract.
type LeeaGovernanceVoteCastWithParamsIterator struct {
	Event *LeeaGovernanceVoteCastWithParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LeeaGovernanceVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeeaGovernanceVoteCastWithParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LeeaGovernanceVoteCastWithParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LeeaGovernanceVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeeaGovernanceVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeeaGovernanceVoteCastWithParams represents a VoteCastWithParams event raised by the LeeaGovernance contract.
type LeeaGovernanceVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_LeeaGovernance *LeeaGovernanceFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*LeeaGovernanceVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _LeeaGovernance.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &LeeaGovernanceVoteCastWithParamsIterator{contract: _LeeaGovernance.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_LeeaGovernance *LeeaGovernanceFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *LeeaGovernanceVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _LeeaGovernance.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeeaGovernanceVoteCastWithParams)
				if err := _LeeaGovernance.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_LeeaGovernance *LeeaGovernanceFilterer) ParseVoteCastWithParams(log types.Log) (*LeeaGovernanceVoteCastWithParams, error) {
	event := new(LeeaGovernanceVoteCastWithParams)
	if err := _LeeaGovernance.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
