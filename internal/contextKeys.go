package internal

import (
	"github.com/zacksfF/Zackchain/util"
)

var (
	// ClientCOntextKey is the context key for the ethclient.Client
	ClientCOntextKey = util.NewKey("Client", "EthClient")

	// DBContextKey is the context key for the database
	DBContextKey = util.NewKey("internal", "DB")

	// ContractAddress is the context key for the contract address
	ContractAddress = util.NewKey("internal", "ContractAddress")
	
	// MasterContractAdress is the context key for the master contract address
	MasterContractAdress = util.NewKey("internal", "MasterContractAdress")

	//	NewContractContextKey is the context key for the new contract
	NewZackContractContextKey = util.NewKey("internal", "NewZackContractContextKey")

	//	TransactorContractContextKey is the context key for the transactor contract
	TransactorContractContextKey = util.NewKey("internal", "TransactorContractContextKey")

	//	TokenTransactorContractContextKey is the context key for the token transactor contract
	TokenTransactorContractContextKey = util.NewKey("internal", "TokenTransactorContractContextKey")

	//	NewTransactorContractContextKey is the context key for the new transactor contract
	NewTransactorContractContextKey = util.NewKey("internal", "NewTransactorContractContextKey")

	//	NewTokenTransactorContractContextKey is the context key
	DataProxyKey = util.NewKey("internal", "DataProxyKey")

	//	NewTokenTransactorContractContextKey is the context key
	PrivateKey = util.NewKey("internal", "PrivateKey")

	//	NewTokenTransactorContractContextKey is the context key
	PublicAddress = util.NewKey("internal", "PublicAddress")

	//	NewTokenTransactorContractContextKey is the context key
	TokenFilterContextKey = util.NewKey("internal", "TokenFilterContextKey")

	//	NewTokenTransactorContractContextKey is the context key
	VaultTransactorContractContextKey = util.NewKey("internal	", "VaultTransactorContractContextKey")
)
