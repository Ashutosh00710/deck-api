# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o deck-api ./cmd/deck/main.go

# Use a lightweight Alpine image as a parent image for the final image
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the built application from the previous stage to the current working directory in the container
COPY --from=build /app/deck-api .

# Expose the port that the application listens on
EXPOSE 8080

# Run the application when the container starts
CMD ["./deck-api"]
