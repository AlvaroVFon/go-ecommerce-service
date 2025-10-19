# Go E-commerce API

## Go E-commerce API

[![Go Version](https://img.shields.io/badge/go-1.24.2-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/docker-%20%20-blue.svg?logo=docker)](https://www.docker.com/)
[![Docker Compose](https://img.shields.io/badge/docker--compose-%20%20-blue.svg?logo=docker)](https://docs.docker.com/compose/)
[![Make](https://img.shields.io/badge/make-%20%20-lightgrey.svg?logo=gnu-make)](#)

Proyecto API REST en Go para una plataforma de e-commerce. Este repositorio contiene la API, la lógica de negocio y utilidades para gestionar productos, usuarios, roles y autenticación.

## Estado actual

El proyecto está avanzado: estructura modular, conexión a PostgreSQL, autenticación JWT, manejo de contraseñas, y endpoints para productos, usuarios y health-check. Principales componentes:

- Estructura: `cmd/`, `internal/`, `pkg/`.
- Configuración: variables por entorno, carga con `godotenv`.
- Base de datos: PostgreSQL con migraciones en `internal/database/migrations` y seeds en `internal/database/seeds`.
- Router HTTP: `chi` (en `internal/*/routes.go`).
- Seguridad: JWT (paquete `github.com/golang-jwt/jwt/v5`) y hashing de contraseñas (`golang.org/x/crypto`).
- Docker: `Dockerfile` y `docker-compose.yml` para desarrollo local.
- Makefile con comandos útiles (arranque, build, tests, docker-up/down, etc.).

## Requisitos

- Go 1.24.2 (ver `go.mod`).
- Docker
- Docker Compose
- make

## Rápido arranque

1. Clona el repositorio:

```bash
git clone https://github.com/AlvaroVFon/go-ecommerce-service.git
cd go-ecommerce-service
```

2. Copia el ejemplo de variables de entorno y ajústalas:

```bash
cp .env.example .env
# editar .env según tu entorno
```

3. Inicia servicios dependientes (Postgres) con Docker Compose:

```bash
make docker-up
```

4. Ejecuta la aplicación en modo desarrollo:

```bash
make run
```

La API quedará disponible en http://localhost:8080 (por defecto).

## Endpoints principales

- Health: GET /health
- Productos: rutas definidas en `internal/products` (crear, listar, obtener, actualizar, eliminar)
- Usuarios: rutas en `internal/users` (registro, login, gestión)
- Roles: `internal/roles` y seeds iniciales en `internal/database/seeds`

Consulta el código en `internal/` para ver handlers y servicios concretos.

## Comandos disponibles (Makefile)

- `make run` - Ejecutar la aplicación en desarrollo
- `make build` - Construir binario
- `make test` - Ejecutar tests
- `make lint` - Ejecutar linters
- `make docker-up` - Levantar contenedores (Postgres)
- `make docker-down` - Parar contenedores
- `make clean` - Limpiar artefactos

## Notas para desarrolladores

- Las migraciones SQL están en `internal/database/migrations`.
- Los seeds iniciales para roles y usuarios están en `internal/database/seeds`.
