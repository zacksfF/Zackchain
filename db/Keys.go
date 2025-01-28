package db

import "strings"

// BalanceKey is the key to store/lookup account balance
const (
	BalanceKey = "eth_balance"

	//currentChallengeKey DB Key
	currentChallengeKey = "current_challenge"
	RequestIDkey        = "current_request_id"
	RequestIDkey0       = "current_request_id0"
	RequestIDkey1       = "current_request_id1"
	RequestIDkey2       = "current_request_id2"
	RequestIDkey3       = "current_request_id3"
	RequestIDkey4       = "current_request_id4"
	DifficultIdKey      = "current_difficulty"
	QueryStringKey      = "current_query_string"
	GranularityKey      = "current_granularity"
	TotalTipKey         = "current_total_tip"
	MiningStatusKey     = "mining_status"

	//GasKey
	GasKEy = "wei_gas_price"

	///Top50
	Top50Key = "top_50_requestIds"

	// TokenBalance
	TokenBalanceKey = "token_balance"

	// Dispute Status
	DisputeStatusKey = "dispute_status"

	//RequestID's are stored with this prefix and the id itself
	//e.g. "qm_2" represents request ID 2
	QueryMetadataPrefix = "qm_"

	//Request values are stored with this prefix plus request id
	QueriedValuePrefix = "qv_"
	LastNewValueKey    = "lastnewvalue"
	LastSubmissionKey  = "last_submission"
	TimeOutKey         = "time_out"
)

var knownKeys map[string]bool

func initKeyLook() {
	knownKeys = make(map[string]bool)
	knownKeys[BalanceKey] = true
	knownKeys[currentChallengeKey] = true
	knownKeys[RequestIDkey] = true
	knownKeys[RequestIDkey0] = true
	knownKeys[RequestIDkey1] = true
	knownKeys[RequestIDkey2] = true
	knownKeys[RequestIDkey3] = true
	knownKeys[RequestIDkey4] = true
	knownKeys[DifficultIdKey] = true
	knownKeys[QueryStringKey] = true
	knownKeys[GranularityKey] = true
	knownKeys[TotalTipKey] = true
	knownKeys[MiningStatusKey] = true
	knownKeys[GasKEy] = true
	knownKeys[Top50Key] = true
	knownKeys[TokenBalanceKey] = true
	knownKeys[DisputeStatusKey] = true
	knownKeys[LastNewValueKey] = true
	knownKeys[LastSubmissionKey] = true
	knownKeys[TimeOutKey] = true
}

func isKnownKey(key string) bool {
	if knownKeys == nil {
		initKeyLook()
	}
	if !knownKeys[key] {
		if !strings.HasPrefix(key, QueryMetadataPrefix) && !strings.HasPrefix(key, QueriedValuePrefix) {
			return false
		}
	}
	return true
}
