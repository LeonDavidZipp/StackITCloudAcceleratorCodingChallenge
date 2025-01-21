goa:
	goa gen github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/design --output ./src/api

start:
	docker-compose up --build

start_local:
	go run ./src/main.go
