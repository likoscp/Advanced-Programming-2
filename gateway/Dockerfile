FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/main.go

FROM alpine:3.19

RUN apk --no-cache upgrade

COPY --from=builder /server /server

EXPOSE 8089

ENTRYPOINT ["/server"]
