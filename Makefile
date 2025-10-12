# Makefile para proyecto E-commerce API (Go)

APP_NAME=ecommerce-api
DB_URL=postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable

# Ejecutar la aplicación
run:
	go run ./cmd/api

# Compilar binario
build:
	go build -o bin/$(APP_NAME) ./cmd/api

# Ejecutar tests con cobertura
test:
	go test ./... -cover

# Ejecutar migraciones
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

# Linter
lint:
	golangci-lint run

# Levantar entorno con Docker
docker-up:
	docker compose up -d

docker-down:
	docker compose down

# Limpiar artefactos de build
clean:
	rm -rf bin
