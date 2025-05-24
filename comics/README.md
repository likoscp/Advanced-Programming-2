go run cmd/main.go
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/comicsdb?sslmode=disable" up
