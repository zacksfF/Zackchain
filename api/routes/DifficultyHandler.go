package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zacksfF/Zackchain/common"
	"github.com/zacksfF/Zackchain/db"
)

// BalanceHandler handles balance request
type DifficultyHandler struct {
}

// Incoming implemntation for handler
func (h *DifficultyHandler) Incoming(ctx context.Context, req *http.Request) (int, string) {
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.DifficultyKey)
	if err != nil {
		log.Printf("problem reading Difficulty from DB: %v\n", err)
		return 500, `{"error": "Could not read Difficulty from DB"}`
	}
	return 200, fmt.Sprintf(`{ "Difficulty": "%s" }`, string(v))
}
