dev:
	docker-compose up -d
	~/go/bin/reflex -s -r '\.go$$' make format run

format:
	go fmt ./...

run:
	go run -race ./cmd/api/main.go

test:
	go test -v -tags testing ./...

test-cov:
	go test -coverprofile=cover.txt ./... && go tool cover -html=cover.txt -o cover.html

build-mocks:
	go get github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golang/mock/mockgen
	~/go/bin/mockgen -source=product_provider/product_provider.go -destination=product_provider/mock/service.go -package=mock
	~/go/bin/mockgen -source=product/product.go -destination=product/mock/service.go -package=mock

dependencies:
	go mod download


up-mongo:
	docker-compose up -d mongodb