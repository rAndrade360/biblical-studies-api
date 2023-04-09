migrate_up:
	export MIGRATIONS_PATH=$(shell pwd)/internal/migrations && go run ./cmd/migration up
migrate_down:
	export MIGRATIONS_PATH=$(shell pwd)/internal/migrations && go run ./cmd/migration down
dev:
	go run cmd/app/main.go