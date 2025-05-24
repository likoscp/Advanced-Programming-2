FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM debian:bookworm-slim
ARG MY_ENV_VAR
ENV MY_ENV_VAR=${MY_ENV_VAR}

RUN apt-get update && apt-get install -y libc6

WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
