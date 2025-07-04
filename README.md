# KART-CHALLENGE

This is a simple e-commerce application built with Go and Gin. It provides APIs for managing products, orders, and users. The application also includes a Redis cache for performance optimization and a PostgreSQL database for data persistence.

## Features

- Product management (CRUD operations)
- Order management (CRUD operations)
- User management (CRUD operations)
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
   Create a `.env` file in the root directory with the following content:
