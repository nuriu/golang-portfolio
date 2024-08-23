FROM golang:1.22.5 AS build-stage
WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -o ./bin/tm ./cmd/api/main.go

EXPOSE 8080

WORKDIR /app/bin
CMD ["./tm"]
