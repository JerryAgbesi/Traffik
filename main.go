package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jerryagbesi/traffik/pkg/loadBalancer"
	"github.com/valyala/fasthttp"
)

func main() {
	var algorithmType = flag.String("algorithm", "random", "Load balancing algorithm to use (random, round-robin)")
	flag.Parse()

	lb := loadBalancer.NewLoadBalancer("configs/servers.json", *algorithmType)

	server := &fasthttp.Server{Handler: loadBalancer.LbRequestHandler(lb), Name: "Traffik Load Balancer"}

	fmt.Printf("Traffik Load Balancer is running on port %s\n", ":4000")

	err := server.ListenAndServe(":4000")
	if err != nil {
		log.Fatalf("Error listening and serving: %v", err)
	}

}
