build:
	go build -o main cmd/api/main.go

test:
	go test -v ./...

test-integration:
	sh scripts/test.sh

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go

docker-build:
	docker build -t uala-microblog-challenge .

docker-run:
	docker run -p 8080:8080 uala-microblog-challenge

docker:
	make docker-build
	make docker-run

prod:
	make docker-build
	make docker-run

local:
	make build
	make test
	make swagger
	make run
