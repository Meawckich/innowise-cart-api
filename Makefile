.PHONY: all

all:
	docker-compose --env-file ./internal/pkg/config/envs/cfg.env up --build

tests:
	go test -v ./...

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -dir cmd/,internal/handler,internal/pkg/db -parseDependency

imports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -local cart-api/ -w .

docker-tests:
	docker-compose --env-file ./internal/pkg/config/envs/cfg.env up --build
	docker-compose exe -T http go test ./...
	docker-compose down

docker-build-test:
	docker build -t api-tests --target=test
 