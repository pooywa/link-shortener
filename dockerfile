# Use a newer Go version
FROM golang:1.22-alpine

# Install sqlite and build dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project
COPY . .

# Initialize go modules and get dependencies
RUN go mod init link-shortener || true
RUN go mod tidy
RUN go mod download

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]