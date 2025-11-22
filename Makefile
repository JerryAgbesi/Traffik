runserver:
	go run cmd/backend/main.go -port 7091

runservers:
	go run cmd/backend/main.go -port 7091 & \
	go run cmd/backend/main.go -port 7092 & \
	go run cmd/backend/main.go -port 7093 & \
	wait

runloadbalancer:
	go run cmd/loadbalancer/main.go -algorithm round-robin
