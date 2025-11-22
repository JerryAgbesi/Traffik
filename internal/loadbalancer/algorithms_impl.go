package loadbalancer

import (
	"math/rand"
	"sync"
	"time"

	"github.com/jerryagbesi/traffik/internal/server"
)

// RandomAlgorithm selects a server randomly.
type RandomAlgorithm struct {
	rnd *rand.Rand
}

func NewRandomAlgorithm() *RandomAlgorithm {
	return &RandomAlgorithm{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *RandomAlgorithm) SelectServer(servers []*server.Server) *server.Server {
	if len(servers) == 0 {
		return nil
	}
	return servers[r.rnd.Intn(len(servers))]
}

func (r *RandomAlgorithm) Name() string {
	return "random"
}

// RoundRobinAlgorithm selects servers in a round-robin fashion.
type RoundRobinAlgorithm struct {
	mu      sync.Mutex
	current int
}

func NewRoundRobinAlgorithm() *RoundRobinAlgorithm {
	return &RoundRobinAlgorithm{
		current: 0,
	}
}

func (r *RoundRobinAlgorithm) SelectServer(servers []*server.Server) *server.Server {
	if len(servers) == 0 {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	server := servers[r.current]
	r.current = (r.current + 1) % len(servers)
	return server
}

func (r *RoundRobinAlgorithm) Name() string {
	return "round-robin"
}
