# Traffik

Traffik is a simple load balancer or proxy server written in Go. I built it to get familiar with Go and understand some load balancing concepts as well.

## Project Structure

The project follows a standard Go layout:

- **`cmd/`**: Application entry points.
  - `cmd/loadbalancer`: The load balancer executable.
  - `cmd/backend`: The backend server executable.
- **`internal/`**: Private application code (algorithms, config, server logic).
- **`configs/`**: Configuration files (e.g., `servers.json`).

## Getting Started

### 1. Start Backend Servers

You can run multiple backend servers on different ports or a single server using the makefile:

```bash
# single server
make runserver

# multiple servers (default 3)
make runservers

```

### 2. Start Load Balancer

Run the load balancer. You can specify the algorithm (`random` or `round-robin`) and the port (default 4000):

```bash
# Run with Round Robin algorithm
go run cmd/loadbalancer/main.go -algorithm round-robin

# Run with Random algorithm (default)
go run cmd/loadbalancer/main.go -algorithm random
```

### 3. Test

Send requests to the load balancer:

```bash
curl http://localhost:4000
```

## Todo
- [x] Create struct and methods for the load balancer 
- [x] Create struct for the backend servers
- [x] Add configuration logic for registering servers
- [x] Implement a proxy handler to redirect traffic from the load balancer to backend servers
- [x] Add load balancer health check for servers
- [ ] Improve the random load balancing algorithm
- [ ] Add support for round robin algorithm
- [ ] Add support for weighted round robin algorithm
- [ ] Allow users to specify config file location with flags
- [ ] Add backend server monitoring (CPU and Memory usage)
