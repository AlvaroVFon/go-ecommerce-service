# Go E-commerce API

[![Go Version](https://img.shields.io/badge/go-1.24.2-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/docker-%20%20-blue.svg?logo=docker)](https://www.docker.com/)
[![Docker Compose](https://img.shields.io/badge/docker--compose-%20%20-blue.svg?logo=docker)](https://docs.docker.com/compose/)
[![Make](https://img.shields.io/badge/make-%20%20-lightgrey.svg?logo=gnu-make)](#)

API REST desarrollada en Go para una plataforma de e-commerce. Este repositorio contiene la API, la lógica de negocio y utilidades para gestionar productos, categorías, carritos, pedidos, usuarios y autenticación.

## Features

- **Gestión de Productos:** CRUD completo para productos.
- **Gestión de Categorías:** CRUD completo para categorías de productos.
- **Gestión de Usuarios:** Registro y obtención de datos de usuario.
- **Autenticación:** Sistema de registro y login basado en JWT.
- **Roles:** Diferenciación entre usuarios normales y administradores.
- **Carrito de Compras:** Lógica para crear y gestionar el carrito de un usuario.
- **Pedidos:** Creación y consulta de pedidos.
- **Salud de la API:** Endpoint de Health-check.

## Requisitos

- Go 1.24.2+
- Docker
- Docker Compose
- Make
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate): Necesario para ejecutar las migraciones de la base de datos.

## Inicio Rápido (Quick Start)

Sigue estos pasos para tener un entorno de desarrollo local funcionando.

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/AlvaroVFon/go-ecommerce-service.git
    cd go-ecommerce-service
    ```

2.  **Configura las variables de entorno:**
    Copia el archivo de ejemplo y ajústalo si es necesario. Los valores por defecto están pensados para el entorno Docker.
    ```bash
    cp .env.example .env
    ```

3.  **Inicia la base de datos:**
    Este comando levantará un contenedor de PostgreSQL con Docker Compose.
    ```bash
    make docker-up
    ```

4.  **Ejecuta las migraciones:**
    Creará todas las tablas necesarias en la base de datos.
    ```bash
    make migrate-up
    ```

5.  **(Opcional) Puebla la base de datos:**
    Este comando inserta datos de prueba (roles, usuarios, productos, etc.) para facilitar el desarrollo.
    ```bash
    make seed
    ```

6.  **Ejecuta la aplicación:**
    ```bash
    make run
    ```

La API estará disponible en `http://localhost:8080` (o el puerto que hayas configurado en tu archivo `.env`).

## Comandos del Makefile

| Comando | Descripción |
| :--- | :--- |
| `make run` | Ejecuta la aplicación en modo desarrollo con hot-reload. |
| `make build` | Compila el binario de la aplicación en la carpeta `bin/`. |
| `make test` | Ejecuta todos los tests del proyecto. |
| `make lint` | Ejecuta el linter `golangci-lint`. |
| `make seed` | Puebla la base de datos con datos de prueba. |
| `make docker-up` | Inicia los servicios de Docker (PostgreSQL). |
| `make docker-down` | Detiene los servicios de Docker. |
| `make migrate-up` | Aplica todas las migraciones pendientes de la base de datos. |
| `make migrate-down` | Revierte la última migración aplicada. |
| `make migrate-new name=<name>` | Crea un nuevo archivo de migración. |
| `make clean` | Elimina los artefactos de compilación. |

## API Endpoints

La URL base de la API es `/api/v1`.

| Método | Ruta | Descripción | Auth | Admin |
| :--- | :--- | :--- | :--- | :--- |
| `GET` | `/health-check` | Comprueba el estado de la API. | No | No |
| `POST` | `/auth/register` | Registra un nuevo usuario. | No | No |
| `POST` | `/auth/login` | Inicia sesión y obtiene un token JWT. | No | No |
| `GET` | `/users/me` | Obtiene los datos del usuario autenticado. | Sí | No |
| `GET` | `/users` | Lista todos los usuarios. | Sí | Sí |
| `GET` | `/users/{userID}` | Obtiene un usuario por su ID. | Sí | Sí |
| `GET` | `/products` | Lista todos los productos. | No | No |
| `GET` | `/products/{productID}` | Obtiene un producto por su ID. | No | No |
| `POST` | `/products` | Crea un nuevo producto. | Sí | Sí |
| `PUT` | `/products/{productID}` | Actualiza un producto existente. | Sí | Sí |
| `DELETE` | `/products/{productID}` | Elimina un producto. | Sí | Sí |
| `GET` | `/categories` | Lista todas las categorías. | No | No |
| `GET` | `/categories/{categoryID}` | Obtiene una categoría por su ID. | No | No |
| `POST` | `/categories` | Crea una nueva categoría. | Sí | Sí |
| `PUT` | `/categories/{categoryID}` | Actualiza una categoría existente. | Sí | Sí |
| `DELETE`| `/categories/{categoryID}`| Elimina una categoría. | Sí | Sí |
| `POST` | `/cart` | Crea un carrito para el usuario. | Sí | No |
| `GET` | `/cart` | Obtiene el carrito del usuario. | Sí | No |
| `POST` | `/cart/items` | Añade un item al carrito. | Sí | No |
| `PATCH` | `/cart/items/{cartItemID}` | Actualiza la cantidad de un item. | Sí | No |
| `DELETE`| `/cart/items/{cartItemID}`| Elimina un item del carrito. | Sí | No |
| `POST` | `/orders` | Crea un pedido a partir del carrito. | Sí | No |
| `GET` | `/orders/{orderID}` | Obtiene un pedido por su ID. | Sí | No |