.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: build
build:
	go build -o main ./cmd/app/main.go