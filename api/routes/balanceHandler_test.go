package routes

import (
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prometheus/tsdb/testutil"
	"github.com/zacksfF/Zackchain/db"
)

type BalResult struct {
	Balance string `json:"balance"`
}

func TestBalanceHandler(t *testing.T) {
	setup(t)
	h := &BalanceHandler{}

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance-rest"))
	if err != nil {
		t.Fatal(err)
	}

	bigBal := big.NewInt(350000)
	raw_bal := hexutil.EncodeBig(bigBal)
	DB.Put(db.BalanceKey, []byte(raw_bal))

	router := NewRouter(DB)
	router.AddRoute("/balance", h)

	http.Handle("/", router)

	go func() { http.ListenAndServe("localhost:5001", nil) }()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:5001/balance") // not sure how to make requests for routes.router
	if err != nil {
		t.Fatalf("Error getting balance from endpoint: %v", err)
	}
	defer resp.Body.Close()
	time.Sleep(30 * time.Second)

	var bal BalResult
	t.Logf("Balance from response: %v", resp.Body)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&bal)
	if !strings.Contains(bal.Balance, "0x") {
		testutil.Ok(t, err)
	} else {
		t.Logf("Retrieved balance from server: %+v\n", bal)
	}
}
