# Makefile para proyecto E-commerce API (Go)

APP_NAME=ecommerce-api
DB_URL=postgres://postgres:postgres@localhost:5433/ecommerce_db?sslmode=disable
MIGRATIONS_DIR = internal/database/migrations

# Ejecutar la aplicación
run:
	go run ./cmd/api

# Ejecutar seeding de datos
seed:
	go run ./cmd/seed
	
# Compilar binario
build:
	go build -o bin/$(APP_NAME) ./cmd/api

# Ejecutar tests con cobertura
test:
	go test ./... -cover

# Ejecutar migraciones
migrate-up:
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DB_URL)" up

migrate-down:
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DB_URL)" down

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

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
