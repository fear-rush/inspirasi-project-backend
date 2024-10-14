# Earthquake API

This project is a Go-based API that provides earthquake data.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Getting Started](#getting-started)
3. [Running Locally](#running-locally)
4. [Environment Variables](#environment-variables)
5. [API Documentation](#api-documentation)
6. [Building the Project](#building-the-project)
7. [Docker Deployment](#docker-deployment)
8. [Testing](#testing)

## Prerequisites

- Go 1.23 or higher
- Docker (optional, for containerized deployment)

## Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/fear-rush/earthquake-api.git
   cd earthquake-api
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Running Locally

To run the project locally:

1. Set up your environment variables (see [Environment Variables](#environment-variables) section).
2. Run the server:
   ```
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080` (or the port specified in your .env file).

## Environment Variables

1. Copy the `.env.example` file to `.env`:
   ```
   cp .env.example .env
   ```

2. Edit the `.env` file and fill in your actual values:
   ```
   BASE_URL=http://api.example.com
   TOKEN_ID=your_token_id
   TOKEN=your_token
   AUTH_COLLECTION_ID=your_auth_collection_id
   USERNAME=your_username
   PASSWORD=your_password
   PROJECT_ID=your_project_id
   COLLECTION_ID=your_collection_id
   ENVIRONMENT=development
   PORT=8080
   ```

## API Documentation

This project uses Swaggo to generate Swagger 2.0 documentation.

To generate or update the Swagger documentation:

1. Install swag CLI:
   ```
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. Generate the documentation:
   ```
   swag init -g cmd/server/main.go
   ```

When running in non-production mode, you can access the Swagger UI at `http://localhost:8080/swagger/index.html`.

## Building the Project

To build the project into a binary:

```
go build -o earthquake-api ./cmd/server
```

This will create an executable named `earthquake-api` in your current directory.

## Docker Deployment

1. Build the Docker image:
   ```
   docker build -t inspirasi-project-backend .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 --env-file .env inspirasi-project-backend or using docker compose up -d 
   (from docker-compose.yaml file)
   ```

Note: Make sure your `.env` file is in the same directory when running the Docker container.

## Testing

To run the integration tests:

```
go test ./test
```

## API Endpoints

- GET `/api/earthquake`: Retrieve earthquake data
- GET `/swagger/*`: Swagger UI (disabled in production)

## Logging

Logs are printed to the console. In a production environment, consider implementing a more robust logging solution.

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
