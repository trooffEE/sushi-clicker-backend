.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: build
build:
	go build -o main ./cmd/app/main.go

create-migration:
	@bash -c 'read -p "Please provide migration name: " name && \
	echo $$name && \
	migrate create -ext sql -dir ./internal/db/migrations/ -seq $$name'