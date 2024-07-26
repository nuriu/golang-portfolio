FROM golang:1.22.5 AS build-stage
WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o ./tm ./cmd/api/main.go

# FROM alpine:latest AS build-release-stage
# WORKDIR /

# COPY --from=build-stage /app/tm /app/tm

EXPOSE 8080

CMD ["/app/tm"]
