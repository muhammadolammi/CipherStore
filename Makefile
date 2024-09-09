build:
	@go build  -o bin/cs . 

run: build
	@./bin/cs


test:
	@go test ./... -v
