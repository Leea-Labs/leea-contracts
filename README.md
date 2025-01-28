# Leea Smart Contracts (Decentralized core)

## Solana based (contracts/solana/)
Prerequisites:
* anchor-cli 0.30.1

##### Project Objects:
* Leea Token aiCO

## EVM based smart contracts (contracts/evm/)
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
* Agent rewards cap/threshold
* system fees % (rewards to validators)
* Agent scoring parameters
* Validator stake/reward/slashing

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

# Deployed at Holesky (EVM)
# Leea Token 0x8cB8AB2a22a882032d277ae29B4c70F60444f95e
# Leea DAO 0x153E8ea256fDC02487882aa48A009D3573C25F99
# Leea Agent Registry 0xe61461139682822a9033A28DDc35377A50edc52e
# Owner 0xDB7B9cd59ebF909D2F29D0278162A17a43fBBb50