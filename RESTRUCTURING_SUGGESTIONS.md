# Project Restructuring Suggestions

## Current Structure Analysis

### Issues Identified

1. **Dual Go Modules**: Two separate Go modules (`github.com/jerryagbesi/traffik` and `traffik-servers`) create confusion
2. **Incomplete Round-Robin**: The round-robin algorithm is not properly implemented (just returns `servers[0]`)
3. **Algorithm Selection**: String-based algorithm selection instead of using interfaces/strategy pattern
4. **Hardcoded Configuration**: Config file path is hardcoded in main.go
5. **Error Handling**: Using `log.Fatalf` instead of proper error propagation
6. **Package Organization**: Mixed concerns in loadBalancer package
7. **Typo**: `ReponseTime` should be `ResponseTime` in server.go
8. **Backend Server Separation**: Backend servers are in a separate module but part of the same project

## Recommended Structure

```
Traffik/
├── cmd/
│   ├── loadbalancer/
│   │   └── main.go              # Load balancer entry point
│   └── backend/
│       └── main.go              # Backend server entry point
├── internal/
│   ├── config/
│   │   └── config.go           # Configuration loading and parsing
│   ├── algorithms/
│   │   ├── algorithm.go        # Algorithm interface
│   │   ├── random.go           # Random algorithm implementation
│   │   ├── roundrobin.go       # Round-robin algorithm implementation
│   │   └── weighted.go         # Weighted round-robin (future)
│   ├── health/
│   │   └── checker.go          # Health checking logic
│   └── proxy/
│       └── proxy.go            # Request proxying logic
├── pkg/
│   ├── server/
│   │   └── server.go           # Server model
│   └── http/
│       └── client.go           # HTTP client utilities
├── configs/
│   └── servers.json            # Server configuration
├── docker/
│   ├── loadbalancer/
│   │   └── Dockerfile
│   └── backend/
│       ├── Dockerfile
│       └── docker-compose.yml
├── scripts/
│   └── requests.sh             # Test script
├── go.mod                      # Single Go module
├── go.sum
└── README.md
```

## Key Improvements

### 1. Single Go Module
- Consolidate into one module: `github.com/jerryagbesi/traffik`
- Use `cmd/` directory for application entry points (standard Go convention)
- Keep shared code in `pkg/` and internal code in `internal/`

### 2. Algorithm Interface Pattern
```go
type Algorithm interface {
    SelectServer(servers []*server.Server) *server.Server
    Name() string
}
```
- Makes it easy to add new algorithms
- Better testability
- Type-safe algorithm selection

### 3. Better Package Organization
- `internal/` for private implementation details
- `pkg/` for reusable packages
- `cmd/` for executables
- Clear separation of concerns

### 4. Configuration Management
- Create a config package to handle configuration loading
- Support command-line flags for config file path
- Environment variable support

### 5. Error Handling
- Replace `log.Fatalf` with proper error returns
- Use structured logging
- Graceful shutdown handling

### 6. Health Checking
- Separate health checking into its own package
- Support configurable health check intervals
- Track server health status

### 7. Proxy Logic Separation
- Extract proxying logic into its own package
- Better testability
- Clearer responsibilities

## Implementation Priority

### High Priority
1. ✅ Fix typo: `ReponseTime` → `ResponseTime`
2. ✅ Consolidate to single Go module
3. ✅ Implement proper round-robin algorithm
4. ✅ Create algorithm interface
5. ✅ Move main.go files to cmd/ directory
6. ✅ Improve error handling

### Medium Priority
1. ✅ Reorganize packages (internal/ structure)
2. ✅ Extract configuration management
3. ✅ Separate health checking logic
4. ✅ Separate proxy logic

### Low Priority
1. ✅ Add structured logging
2. ✅ Add graceful shutdown
3. ✅ Improve Docker organization
4. ✅ Add more comprehensive tests

## Migration Steps

1. **Create new directory structure**
2. **Move and refactor code**:
   - Move `main.go` → `cmd/loadbalancer/main.go`
   - Move `backendServers/main.go` → `cmd/backend/main.go`
   - Refactor loadBalancer package into internal packages
3. **Update imports** throughout the codebase
4. **Fix algorithm implementation**
5. **Update Docker files** to new structure
6. **Test thoroughly**

## Additional Recommendations

1. **Add Tests**: Create unit tests for algorithms, health checks, and proxy logic
2. **Add Metrics**: Consider adding Prometheus metrics for monitoring
3. **Add Middleware**: Support for request/response middleware
4. **Configuration Validation**: Validate server configuration on startup
5. **Documentation**: Add godoc comments to exported functions
6. **CI/CD**: Add GitHub Actions for testing and building

