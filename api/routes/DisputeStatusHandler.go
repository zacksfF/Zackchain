package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zacksfF/Zackchain/common"
	"github.com/zacksfF/Zackchain/db"
)

// BalanceHandler handles balance requests
type DisputeStatusHandler struct {
}

// BalanceHandler implematation for handler
func (h *DisputeStatusHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		log.Printf("Problem reading DisputeStatus from DB: %v\n", err)
		return 500, `{"error": "Could not read DisputeStatus from DB"}`
	}
	return 200, fmt.Sprintf(`{ "DisputeStatus": "%s" }`, string(v))
}
