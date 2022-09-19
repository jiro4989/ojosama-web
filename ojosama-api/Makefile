ojosama-api: go.* *.go
	make test
	go vet .
	go fmt .
	go build

.PHONY: test
test:
	go test -cover ./...
