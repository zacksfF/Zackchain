package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zacksfF/Zackchain/common"
	"github.com/zacksfF/Zackchain/db"
)

//BalanceHandler handles balance requests
type CurrentChallengeHandler struct{}

//Incoming implematation for handler 
func (h *CurrentChallengeHandler) Incoming(ctx context.Context, req *http.Request) (int, string){
	DB := ctx.Value(common.DBContextKey).(db.DB)
	v, err := DB.Get(db.CurrentChallengeKey)
	if err != nil{
		log.Printf("Problem reading currentChallenge from DB: %v\n", err)
		return 500, `{"error":"Could not read currentChallenge from DB"}`
	}
	return 200, fmt.Sprintf(`{"currentChallenge":"%s"}`, string(v))
}