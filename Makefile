lint:
	@golint ./...
	
test:
	@go test -v ./...

build: 
	@go build -i -v ./...

install:
	@go install -v ./cmd/payments-webhook

run-local:
	@docker-compose -f docker-compose.yml up --build --force-recreate
