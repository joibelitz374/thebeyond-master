proto-gen:
	buf lint
	buf build
	buf generate

run-db-dev:
	docker-compose -f ./deployments/dev.compose.yaml up -d

run-api-dev:
	go run cmd/api/main.go

run-bot-dev:
	go run cmd/bot/main.go

fix-migrations:
	sqlfluff fix ./migrations/* --dialect postgres