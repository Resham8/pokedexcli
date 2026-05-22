FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o pokedexcli

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/pokedexcli .

CMD ["./pokedexcli"]