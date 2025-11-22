package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jerryagbesi/traffik/internal/loadbalancer"
	"github.com/valyala/fasthttp"
)

func main() {
	var (
		algorithmType = flag.String("algorithm", "random", "Load balancing algorithm to use (random, round-robin)")
		configFile    = flag.String("config", "configs/servers.json", "Path to the server configuration file")
		port          = flag.String("port", "4000", "Port to run the load balancer on")
	)
	flag.Parse()

	var algo loadbalancer.Algorithm
	switch *algorithmType {
	case "random":
		algo = loadbalancer.NewRandomAlgorithm()
	case "round-robin":
		algo = loadbalancer.NewRoundRobinAlgorithm()
	default:
		log.Fatalf("Unknown algorithm type: %s", *algorithmType)
	}

	lb, err := loadbalancer.NewLoadBalancer(*configFile, algo)
	if err != nil {
		log.Fatalf("Error creating load balancer: %v", err)
	}

	server := &fasthttp.Server{
		Handler: loadbalancer.LbRequestHandler(lb),
		Name:    "Traffik Load Balancer",
	}

	fmt.Printf("Traffik Load Balancer is running on port :%s using %s algorithm\n", *port, algo.Name())

	err = server.ListenAndServe(":" + *port)
	if err != nil {
		log.Fatalf("Error listening and serving: %v", err)
	}
}
