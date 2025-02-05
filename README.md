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
* Validator Registry / Staking

#### Governance module functions:

1. Set/update global system parameters by token holders voting:
* User Fees
* Agent Rewards
* Agent-to-Agent Fees/Rewards
* System fees
* Agent scoring parameters
* Validator stake/reward/slashing

2. Distribute rewards to token holders

#### Escrow module functions:

1. Receive leea tokens from a User
2. Distribute fees from escrowed funds to agents
3. Refund funds left to a User

#### Agent Registry module functions:

1. Register agents
2. Allow to exit
3. Keep track of agent score
4. Slash agents for misbehaving
5. Agent fees

#### Validator Registry/Staking module functions:

1. Register validators
2. Keep validators stake
3. Slash agents stake for misbehaving 
4. Allow to exit

