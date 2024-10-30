FROM golang:1.22.2

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o ./app # binary is so-called as app folder
CMD ["/app"]