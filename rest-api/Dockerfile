FROM golang:1.22.0 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin ./cmd/app/main.go

FROM scratch AS deploy
WORKDIR /

COPY --from=build /bin /bin

EXPOSE 8080

ENTRYPOINT ["/bin"]