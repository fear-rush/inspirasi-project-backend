# Stage 1: Build stage using Golang 1.23
FROM golang:1.23 as build

# Set the working directory in the container
WORKDIR /app

# Copy only go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project except the files you want to ignore
# Use a .dockerignore file to ignore docs, .env, .env.example, .gitignore, and docker-compose.yml
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./cmd/server/main.go

# Stage 2: Runtime stage using distroless image
FROM gcr.io/distroless/static-debian12

# Copy the compiled binary from the build stage to the runtime image
COPY --from=build /go/bin/app /app/main

# Define the command to run the app
CMD ["/app/main"]
