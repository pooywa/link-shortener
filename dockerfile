# Use the official Go image as the base image
FROM golang:1.21-alpine

# Install sqlite and build dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]