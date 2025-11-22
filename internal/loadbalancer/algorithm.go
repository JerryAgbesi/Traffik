package loadbalancer

import "github.com/jerryagbesi/traffik/internal/server"

// Algorithm defines the interface for load balancing algorithms.
type Algorithm interface {
	// SelectServer picks a backend server based on the algorithm's logic.
	SelectServer(servers []*server.Server) *server.Server
	// Name returns the name of the algorithm.
	Name() string
}
