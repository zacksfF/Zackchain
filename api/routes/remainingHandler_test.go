package routes

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/math"
// 	zapCommon "github.com/zacksfF/Zackchain/common"
// 	"github.com/zacksfF/Zackchain/config"
// 	"github.com/zacksfF/Zackchain/contracts"
// 	"github.com/zacksfF/Zackchain/db"
// 	"github.com/zacksfF/Zackchain/util"
// )

// var DB db.DB
// var err error
// var ctx context.Context
// var opts *rpc.MockOptions
// var cfg *config.Config
// var currentChallenge [32]byte
// var requestID *big.Int
// var difficulty *big.Int
// var queryString string
// var granularity *big.Int
// var totalTip *big.Int
// var myStatus bool

// func setup(t *testing.T) {
// 	err := config.ParseConfig("../../config.json")
// 	if err != nil {
// 		fmt.Errorf("Can't parse config for test.")
// 	}
// 	path := "../../testConfig.json"
// 	err = util.ParseLoggingConfig(path)
// 	if err != nil {
// 		fmt.Errorf("Can't parse logging config for test.")
// 	}

// 	prep(t)
// }

// func prep(t *testing.T) {
// 	cfg = config.GetConfig()
// 	DB, err = db.Open(cfg.DBFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	startBal := big.NewInt(356000)

// 	hash := math.PaddedBigBytes(big.NewInt(256), 32)
// 	var b32 [32]byte
// 	for i, v := range hash {
// 		b32[i] = v
// 	}
// 	queryStr := "json(https://coinbase.com)"
// 	chal := &rpc.CurrentChallenge{
// 		ChallengeHash: b32,
// 		RequestID:     big.NewInt(1),
// 		Difficulty:    big.NewInt(500),
// 		QueryString:   queryStr,
// 		Granularity:   big.NewInt(1000),
// 		Tip:           big.NewInt(0),
// 	}
// 	opts = &rpc.MockOptions{
// 		ETHBalance:       startBal,
// 		Nonce:            1,
// 		GasPrice:         big.NewInt(700000000),
// 		TokenBalance:     big.NewInt(0),
// 		MiningStatus:     true,
// 		Top50Requests:    []*big.Int{},
// 		CurrentChallenge: chal,
// 		DisputeStatus:    big.NewInt(1),
// 	}

// 	client := rpc.NewMockClientWithValues(opts)

// 	ctx = context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
// 	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)

// 	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)
// 	if masterInstance == nil {
// 		contractAddress := common.HexToAddress(cfg.ContractAddress)
// 		// TODO create error state flag for mock client
// 		masterInstance, err = contracts.NewZapMaster(contractAddress, client)
// 		if err != nil {
// 			t.Errorf("Problem creating zap master instance: %v\n", err)
// 		}
// 		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
// 	}

// 	_fromAddress := cfg.PublicAddress

// 	//convert to address
// 	fromAddress := common.HexToAddress(_fromAddress)

// 	instance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)
// 	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err = instance.GetCurrentVariables(nil)
// 	if err != nil {
// 		fmt.Println("Current Variables Retrieval Error")
// 	}

// 	myStatus, err = instance.DidMine(nil, currentChallenge, fromAddress)
// 	if err != nil {
// 		t.Fatalf("Could not retrieve mining status: %v", err)
// 	}

// }
