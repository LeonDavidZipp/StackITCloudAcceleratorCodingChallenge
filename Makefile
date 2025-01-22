goa:
	goa gen github.com/LeonDavidZipp/StackITCloudAcceleratorCodingChallenge/src/api/design --output ./src/api

start:
	docker compose up --build

start_local:
	go run ./src/main.go

test:
	docker compose up --build -d
	postman collection run StackIT.postman_collection.json
	docker compose down
