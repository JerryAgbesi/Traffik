# Traffik
Traffik is a simple load balancer or proxy server written in Go. I built it to get familiar with Go and understand
some load balancing concepts as well. 

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


