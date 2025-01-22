# Use Go 1.23
FROM golang:1.23-rc-alpine3.19

# Install sqlite and build dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project
COPY . .

# Set GOTOOLCHAIN to allow using newer versions
ENV GOTOOLCHAIN=auto

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