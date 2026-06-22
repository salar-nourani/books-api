build:
	go build -o bin/books-api ./cmd/api/main.go

test:
	go test ./...

run:
	go run ./cmd/api/main.go

clean:
	rm -rf bin/

