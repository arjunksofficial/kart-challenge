# KART-CHALLENGE

This is a simple e-commerce application built with Go and Gin. It provides APIs for managing products, orders, and users. The application also includes a Redis cache for performance optimization and a PostgreSQL database for data persistence.

## Features

- Product management (CRUD operations)
- Order management (CRUD operations)
- Redis caching for performance
- PostgreSQL database for data persistence
- Promocode validation logic
- Swagger documentation for APIs
- Unit tests for core functionalities
- Docker support for containerization
- Makefile for build and run commands

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker (optional, for running Redis and PostgreSQL)
- Make (optional, for using the Makefile)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/kart-challenge.git
   cd kart-challenge
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up environment variables:
   Create a `.env` file in the root directory.
   Refer to the `.env.example` file for required variables.
4. Start PostgreSQL and Redis:
   If you have Docker installed, you can run:
   ```bash
   docker-compose -f build/docker/docker-compose.yaml up -d postgres redis
   ```
   This will start PostgreSQL and Redis containers.
5. Run Promo Code Importer

   Download promo code files to ./promos folder and
   add the paths like below to .env file

   ```
   FILE1_SOURCE=~/Documents/arjun/projects/kart-challenge/.promos/file1.txt
   FILE2_SOURCE=~/Documents/arjun/projects/kart-challenge/.promos/file2.txt
   FILE3_SOURCE=~/Documents/arjun/projects/kart-challenge/.promos/file3.txt
   ```

   If you have a promo code file, run the importer:

   ```bash
   go run cmd/promoimporter/main.go
   ```

   This will read the promo codes from the specified files and import them into the database. This job also make sure only valid promocodes are added to redis

6. Run the migrations for PostgreSQL:

   ```bash
   go run cmd/migrator/main.go
   ```

   This will create the necessary database tables. and create sample products

7. Run the api server
   ```
   go run cmd/apiserver/main.go
   ```
   Check health using command
   ```
   curl http://localhost:9000/health
   ```

To test Golang Unit Test Coverage
Run

```
make test-coverage
```

Notes

- Used mockery to generate mocks
  To generate mocks

```
make generatemocks
```

Some improvements which were left because of time constrains.

- Improve unit test coverage
- Add Integration and load tests
- Create benchmarking tests to measure performance
