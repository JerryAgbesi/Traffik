package loadBalancer

import (
	"log"
	"math/rand"

	"github.com/jerryagbesi/traffik/pkg/http"
	"github.com/jerryagbesi/traffik/pkg/server"
)


func (lb *LoadBalancer) getRandomServer() *server.Server {
	if len(lb.servers) == 0 {
		return nil
	}

	maxRetries := len(lb.servers)

	for i:=0; i < maxRetries ; i++ {
		selectedServer := lb.servers[rand.Intn(len(lb.servers))]

		if http.NewHttpClient().IsHostAlive(selectedServer.URL.String()) {
			log.Printf("server: %s is alive\n", selectedServer.URL.String())
			return selectedServer
		}
		
		log.Printf("server: %s did not respond \n", selectedServer.URL.String())

		}
		
	return nil

}