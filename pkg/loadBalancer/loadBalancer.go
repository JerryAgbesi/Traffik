package loadBalancer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/jerryagbesi/traffik/pkg/server"
	"github.com/valyala/fasthttp"
)

type LoadBalancer struct {
	servers       []*server.Server
	algorithmType string
}

func NewLoadBalancer(filename, algorithmType string) *LoadBalancer {
	lb := &LoadBalancer{algorithmType: algorithmType}
	lb.readJson(filename)
	return lb
}

func (lb *LoadBalancer) readJson(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = json.Unmarshal(content, &lb.servers)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
}

// This is a function I would like to explore using in the future for slowly draining backend servers
// when they don't report a healthy status
func drainBackend(server *server.Server) {
	for i := server.ActiveConns; i > 0; i-- {
		time.Sleep(100 * time.Millisecond)
	}
}

func releaseConnection(server *server.Server) {
	atomic.AddInt32(&server.ActiveConns, -1)
	
}

func LbRequestHandler(lb *LoadBalancer) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var backendServer *server.Server

		switch lb.algorithmType {
		case "random":
			backendServer = lb.getRandomServer()
		case "round-robin":
			backendServer = lb.servers[0]
		default:
			log.Fatalf("Unknown algorithm type: %s", lb.algorithmType)
		}

		if backendServer == nil {
			ctx.Error("No backend servers available", fasthttp.StatusServiceUnavailable)
			return
		}

		startTime := time.Now()

		atomic.AddInt32(&backendServer.ActiveConns, 1)
		defer releaseConnection(backendServer)

		// Handling the proxying of the request to the selected backend server
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		backendURI := fmt.Sprintf("%s%s", backendServer.URL.String(), ctx.Path())
		req.SetRequestURI(backendURI)

		// Making sure all headers,method and body from the original request is set for the proxied request
		req.Header.SetMethodBytes(ctx.Method())
		ctx.Request.Header.CopyTo(&req.Header)
		req.Header.SetHost(backendServer.URL.Host)
		req.SetBody(ctx.PostBody())

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := fasthttp.Do(req, resp)
		if err != nil {
			log.Fatalf("Error proxying request: %v", err)
			ctx.Error("Failed to forward request to backend", fasthttp.StatusInternalServerError)
		}

		// Copy and send over the original response from the backend server
		ctx.Response.SetStatusCode(resp.StatusCode())
		ctx.Write(resp.Body())
		resp.Header.CopyTo(&ctx.Response.Header)

		elapsedTime := time.Since(startTime)

		backendServer.ReponseTime = elapsedTime

		log.Printf("Request to %s took %v", backendServer.URL.String(), elapsedTime)
		log.Printf("Active connections to %s: %d", backendServer.URL.String(), backendServer.ActiveConns)	
	}
}
