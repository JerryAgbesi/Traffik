package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	
	"github.com/valyala/fasthttp"
)

var reqCount uint32

type response struct {
	Message    string `json:"message"`
	Hostname   string `json:"hostname"`
	ReqCounter uint32 `json:"req_counter"`
	// Headers    map[string]string `json:"headers"`	
}

func main() {
	server := &fasthttp.Server{Handler: requestHandler, Name: "Traffik Backend Servers"}

	err := server.ListenAndServe(":8080")
	if err != nil {
		log.Fatalf("Error listening and serving: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	hostname := os.Getenv("hostname")
	res := response{
		Message:    fmt.Sprintf("A JSON response from %s counter %d", hostname, atomic.LoadUint32(&reqCount)),
		Hostname:   hostname,
		ReqCounter: atomic.LoadUint32(&reqCount),
	}

	ctx.SetContentType("application/json")
	ctxBody, _ := json.Marshal(res)
	ctx.SetBody(ctxBody)
	atomic.AddUint32(&reqCount, 1)
}
