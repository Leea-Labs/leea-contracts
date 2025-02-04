### Leea Smart Contracts (Decentralized core)

Prerequisites:

* solc 0.8.22
* abigen 1.14.12-stable

```sh
brew tap ethereum/ethereum
brew install solidity
```

Compile contracts:

```sh
cd contracts
npm install
./gen.sh
```

##### Project Objects:
* Governance (DAO) 
* Leea Token
* Escrow
* Agent Registry 
* Validator Registry 
* Validator Staking

#### Governance module functions:

1. Set system parameters by token holders voting:
* User Fees
* Agent Rewards
* Validator stake
* Validator rewards

