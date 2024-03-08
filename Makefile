dev:
	docker-compose up -d
	~/go/bin/reflex -s -r '\.go$$' make format run

format:
	go fmt ./...

run:
	go run -race ./cmd/api/main.go

test:
	go test -v -tags testing ./...

dependencies:
	go mod download

test-cov:
	go test -coverprofile=cover.txt ./... && go tool cover -html=cover.txt -o cover.html

up-mongo:
	docker-compose up -d mongodb

build-mocks:
	go get github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golang/mock/mockgen
	~/go/bin/mockgen -source=user/user.go -destination=user/mock/service.go -package=mock

build:
	go build cmd/api/main.go