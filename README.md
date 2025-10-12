# Go E-commerce API

This project is a Go-based API for an e-commerce platform. It is designed to be a microservice that will eventually interact with other services, such as an authentication service.

## Current Status

The project is in its initial setup phase. The following has been configured:

- **Project Structure:** A standard Go project layout with `cmd`, `internal`, and `pkg` directories.
- **Configuration:** The application is configured using environment variables, with an `.env.example` file provided as a template. The `godotenv` library is used to load these variables.
- **Database:** The project is set up to connect to a PostgreSQL database. The `pq` driver is used for the connection.
- **HTTP Server:** An HTTP server is set up using the `chi` router.
- **Health Check:** A basic health check endpoint is available at `/health` to monitor the service's status.
- **Docker:** The project includes a `Dockerfile` for building a production-ready container and a `docker-compose.yml` file for setting up a local development environment with a PostgreSQL database.
- **Makefile:** A `Makefile` is provided with commands for common development tasks.

## Getting Started

### Prerequisites

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [make](https://www.gnu.org/software/make/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/go-ecommerce-service.git
    cd go-ecommerce-service
    ```

2.  **Create a `.env` file:**

    Copy the `.env.example` file to a new file named `.env` and update the variables as needed.

    ```bash
    cp .env.example .env
    ```

3.  **Start the database:**

    Use Docker Compose to start the PostgreSQL database service.

    ```bash
    make docker-up
    ```

4.  **Run the application:**

    ```bash
    make run
    ```

The API will be running at `http://localhost:8080`.

## Available Commands

The following commands are available in the `Makefile`:

- `run`: Run the application.
- `build`: Build the application binary.
- `test`: Run the tests.
- `lint`: Lint the code.
- `docker-up`: Start the Docker containers for the development environment.
- `docker-down`: Stop the Docker containers.
- `clean`: Clean up build artifacts.
