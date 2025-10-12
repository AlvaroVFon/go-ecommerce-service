# AI Agent Context for Go E-commerce API

This document provides context for AI agents interacting with the Go E-commerce API project.

## Project Overview

This is a microservice-based e-commerce API written in Go. The primary goal of this service is to handle e-commerce functionalities like products, orders, and payments. Authentication is handled by a separate, external authentication microservice.

## Key Technologies

- **Language:** Go
- **Web Framework:** chi
- **Database:** PostgreSQL
- **Configuration:** Environment variables (managed with `godotenv`)
- **Containerization:** Docker and Docker Compose

## Project Structure

- `cmd/api/main.go`: The main entry point for the application.
- `internal/`: Contains the core application logic, separated by domain (e.g., `product`, `order`).
- `pkg/`: Contains shared libraries and utilities.
- `api/`: Contains OpenAPI/Swagger specifications.

## Development Workflow

- Use the `Makefile` for common tasks like running, building, and testing the application.
- The application is configured through environment variables. A `.env.example` file is provided as a template.
- The local development environment can be started with `make docker-up`.
