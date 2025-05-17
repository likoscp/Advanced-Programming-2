FROM golang:1.23

WORKDIR /app

RUN apt-get update && apt-get install -y \
    sqlite3 \
    libsqlite3-dev \
    && go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN go build -o main ./cmd/app/main.go

EXPOSE 8080

CMD ["./main"]