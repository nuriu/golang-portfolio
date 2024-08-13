FROM golang:1.22.5 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -o ./tm ./cmd/api/main.go

EXPOSE 8080

CMD ["/app/tm"]
