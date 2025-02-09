package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/zacksfF/Zackchain/common"
	"github.com/zacksfF/Zackchain/db"
)

// Handler handle incoming http requests and prioducts a response code and JSpayload string
type Handler interface {
	Incoming(ctx context.Context, req *http.Request) (int, string)
}

// Route is a struct that holds the route and the handler for that route
type Router struct {
	DB     db.DB
	routes map[string]Handler
}

// NewRouter creates a new router
func NewRouter(db db.DB) *Router {
	return &Router{
		DB:     db,
		routes: make(map[string]Handler),
	}
}

// Default http handler callback which willl route to appropriate handler internally
func (r *Router) ServeHTTP(w http.ResponseWriter, req http.Request) {
	w.Header().Add("Content-Type", "application/json")
	url := strings.ToLower(req.URL.String())
	h := r.routes[url]
	if h == nil {
		w.WriteHeader(404)
		return
	}
	log.Printf("Incoming request", url)
	if e := recover(); e != nil {
		fmt.Printf("Problem witj controller %s: %v\n", url, e)
		fmt.Fprintf(w, "Problem handling request")
		return
	}
	ctx := context.WithValue(context.Background(), common.DBContextKey, r.DB)
	code, payload := h.Incoming(ctx, &req)
	w.WriteHeader(code)
	fmt.Fprintf(w, payload)
}

// AddRoute adds handler for the given url
func (r *Router) AddRoute(path string, handler Handler) {
	r.routes[strings.ToLower(path)] = handler
}
