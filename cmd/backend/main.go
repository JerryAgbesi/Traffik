package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync/atomic"

	"github.com/valyala/fasthttp"
)

var reqCount uint32

type response struct {
	Message    string `json:"message"`
	// Hostname   string `json:"hostname"`
	ReqCounter uint32 `json:"req_counter"`
	// Headers    map[string]string `json:"headers"`
}

func main() {
	port := flag.String("port", "8080", "Port to run the backend server on")
	flag.Parse()

	server := &fasthttp.Server{Handler: requestHandler, Name: "Traffik Backend Server"}

	fmt.Printf("Backend server running on port :%s\n", *port)
	err := server.ListenAndServe(":" + *port)
	if err != nil {
		log.Fatalf("Error listening and serving: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	// hostname := os.Getenv("hostname")
	res := response{
		Message:    fmt.Sprintf("A JSON response from counter %d", atomic.LoadUint32(&reqCount)),
		ReqCounter: atomic.LoadUint32(&reqCount),
	}

	ctx.SetContentType("application/json")
	ctxBody, _ := json.Marshal(res)
	ctx.SetBody(ctxBody)
	atomic.AddUint32(&reqCount, 1)
}
