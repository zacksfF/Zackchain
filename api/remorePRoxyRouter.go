package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zacksfF/Zackchain/common"
	"github.com/zacksfF/Zackchain/db"
	"github.com/zacksfF/Zackchain/util"
)

// RemoteProxyRouter handles incoming http requests
type RemoteProxyRouter struct {
	dataProxy db.DataServerProxy
	log       *util.Logger
}

// CreateProxyRouter creates a remote proxy instance
func CreateRemoteProxy(ctx context.Context) (*RemoteProxyRouter, error) {
	proxy := ctx.Value(common.DataProxyKey).(db.DataServerProxy)
	return &RemoteProxyRouter{dataProxy: proxy, log: util.NewLogger("rest", "RemoteProxyRouter")}, nil
}

// Default http handler callback which will route to appropriate handler internally
func (r *RemoteProxyRouter) ServeHTTP(w http.ResponseWriter, req http.Request) {
	w.Header().Add("Content-Type", "application/octet-stream")

	if e := recover(); e != nil {
		fmt.Printf("Problem with proxied data request: %v\v", e)
		fmt.Fprintf(w, "Cannot serve request")
	}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.log.Error("Problem reading request data: %v")
		fmt.Fprintf(w, "Could not read request data")
		return
	}

	r.log.Info("Getting request with %d bytes of data", len(data))
	outData, err := r.dataProxy.IncomingRequest(data)

	if err != nil {
		r.log.Error("Problem handling incoming request data: %v", err)
		fmt.Fprintf(w, "Could not handle request")
		return
	}
	r.log.Info("Produced result with %d bytes of data", len(outData))
	w.WriteHeader(200)
	w.Write(outData)
}
