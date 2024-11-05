FROM golang:1.23.1

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/app/main.go
CMD ["./main"]