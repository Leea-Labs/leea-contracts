#!/bin/bash
solc --abi LeeaGovernance.sol -o compiled
abigen --abi=./compiled/LeeaGovernance.abi --pkg=governance --out=./compiled/LeeaGovernance.go