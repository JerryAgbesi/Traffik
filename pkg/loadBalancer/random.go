package loadBalancer

import(
	"math/rand"

	"github.com/jerryagbesi/traffik/pkg/server"
)


func (lb *LoadBalancer) getRandomServer() *server.Server {
	return lb.servers[rand.Intn(len(lb.servers))]
}