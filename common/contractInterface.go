package common

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zacksfF/Zackchain/db"
)

// TransactionGeneratorFN is a callback function that a TransactionSubmitter uses to actually invoke
// a contract and generate a transaction
type TransactionGeneratorFN func(ctx context.Context, contract ContractInterface) (*types.Transaction, error)

// TransactionGeneratorFN is a callback function that a TransactionSubmitter uses to actually invoke
// a contract and generate a transaction
type ContractInterface interface {
	AddTip(ctx context.Context, requestID big.Int) (*types.Transaction, error)
	SubmitSolution(solutionID string, requestID big.Int, value big.Int) (*types.Transaction, error)
	DidMine(challenge [32]byte) (bool, error)
}

// TransactionSubmitter is an abstraction for something that will callback a generator fn
// with a contract able to submit and generate transactions. This, like the ContractInterface,
// is an abstraction mainly so we can test isolated functionality.
type TransactionSubmitter interface {
	//PrepareTransaction creates a ContractInterface and sends it to the generatorFN. The ctxName is
	//primarily for logging under which context the transaction is being prepared.
	PrepareTransaction(ctx context.Context, proxy db.DataServerProxy, ctxName string, factoryFn TransactionGeneratorFN) error
}
