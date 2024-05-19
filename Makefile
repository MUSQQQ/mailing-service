test:
	@go test -v ./...

build: 
	@go build -v -o ./bin ./...

install:
	@go install -v ./cmd/mailing-service

run-local:
	@docker-compose -f docker-compose.yml up --build --force-recreate
