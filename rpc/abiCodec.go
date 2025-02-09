package rpc

import (
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zacksfF/Zackchain/contracts"
	"github.com/zacksfF/Zackchain/util"
)

var ABICodecLog = util.NewLogger("rpc", "ABICodec")

// ABICodec holds abi definitions for encoding/decoding contract methods and events
type ABICodec struct {
	abiStruct abi.ABI
	methods   map[string]*abi.Method
	Events    map[string]*abi.Event
}

// BuildCodec constructs a merged abi structure representing all methods/events for Zack contracts. This is primarily
// used for mock encoding/decoding parameters but could also be used for manual RPC operations that do not
func BuildCodec() (*ABICodec, error) {
	all := []string{
		contracts.ZapDisputeABI,
		contracts.ZapGettersABI,
		contracts.ZapGettersLibraryABI,
		contracts.ZapStakeABI,
	}

	parsed := make([]interface{}, 0)
	for _, abi := range all {
		var f interface{}
		if err := json.Unmarshal([]byte(abi), &f); err != nil {
			return nil, err
		}
		asList := f.([]interface{})
		for _, parsedABI := range asList {
			parsed = append(parsed, parsedABI)
		}
	}

	j, err := json.Marshal(parsed)
	if err != nil {
		return nil, err
	}
	abiStruct, err := abi.JSON(strings.NewReader(string(j)))
	if err != nil {
		return nil, err
	}
	methodMap := make(map[string]*abi.Method)
	eventMap := make(map[string]*abi.Event)
	for _, a := range abiStruct.Methods {
		sig := hexutil.Encode(a.ID)
		abiCodecLog.Debug("Mapping method sig: %s to method: %s", sig, a.Name)
		methodMap[sig] = &abi.Method{Name: a.Name, Constant: a.Constant, Inputs: a.Inputs, Outputs: a.Outputs}
	}
	for _, e := range abiStruct.Events {
		sig := hexutil.Encode(e.ID.Bytes())
		abiCodecLog.Debug("Mapping event sig: %s to event %s", sig, e.Name)
		eventMap[sig] = &abi.Event{Name: e.Name, Anonymous: e.Anonymous, Inputs: e.Inputs}
	}

	return &ABICodec{abiStruct, methodMap, eventMap}, nil
}

