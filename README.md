# Link Shortener

A simple URL shortener service built with Go, using:
- Gin framework for HTTP routing
- SQLite for storage
- Clean architecture principles

## Setup

1. Clone the repository
2. Run \`go mod tidy\` to install dependencies
3. Run \`go run main.go\` to start the server

## API Endpoints

- POST /shorten - Shorten a URL
  \`\`\`json
  {
    \"long_url\": \"https://example.com\"
  }
  \`\`\`

- GET /:code - Redirect to original URL