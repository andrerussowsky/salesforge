FROM golang:1.15 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o salesforge-api cmd/salesforge-api/main.go

FROM debian:buster AS final

WORKDIR /app

COPY --from=builder /app/salesforge-api .

EXPOSE 8080

CMD ["./salesforge-api"]