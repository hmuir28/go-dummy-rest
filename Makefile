build:
	@go build -o ./bin/go-todo

run: build
	@./bin/go-todo

test:
	go test -v ./...
