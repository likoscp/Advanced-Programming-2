FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y libc6

WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
