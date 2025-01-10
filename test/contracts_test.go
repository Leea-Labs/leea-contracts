package contracts

import (
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	storage_contract "github.com/Leea-Labs/leea-contracts/contracts/artifacts/storage"
)

func TestStorageContract(t *testing.T) {
	sim, auth, err := SetupBackend()
	require.NoError(t, err)
	_, tx, store, err := storage_contract.DeployStorage(auth, sim.Client())
	require.NoError(t, err)
	sim.Commit()
	tx, err = store.Store(auth, big.NewInt(420))
	require.NoError(t, err)
	fmt.Printf("State update pending: 0x%x\n", tx.Hash())
	sim.Commit()
	val, err := store.Retrieve(nil)
	require.Equal(t, big.NewInt(420).Uint64(), val.Uint64())
}

func SetupBackend() (*simulated.Backend, *bind.TransactOpts, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to generate key: %v", err)
	}

	// Since we are using a simulated backend, we will get the chain ID
	// from the same place that the simulated backend gets it.
	chainID := params.AllDevChainProtocolChanges.ChainID

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Failed to make transactor: %v", err)
	}

	sim := simulated.NewBackend(map[common.Address]types.Account{
		auth.From: {Balance: big.NewInt(9e18)},
	})
	return sim, auth, nil
}
