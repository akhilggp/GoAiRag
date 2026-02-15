run:
	go run ./cmd/api

test:
	go test ./...

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run ./...

tidy:
	go mod tidy
