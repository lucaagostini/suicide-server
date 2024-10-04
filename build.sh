#!/bin/bash

set -e

# Colors for output
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Application name
APP_NAME="lucaagostini/suicide-server"

echo -e "${GREEN}Building $APP_NAME...${NC}"

# Build Go binary
echo "Building Go binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o target/main -ldflags="-w -s" .

# Build Docker image
echo "Building Docker image..."
docker build --platform="linux/amd64" -t $APP_NAME .

echo -e "${GREEN}Build complete!${NC}"

# Print image size
echo "Docker image details:"
docker images $APP_NAME --format "{{.Repository}}:{{.Tag}} - {{.Size}}"

echo -e "${GREEN}To run the container, use:${NC}"
echo "docker run --rm -p 8080:8080 -e SUICIDE_HEALTH_AFTER_SECONDS=60 $APP_NAME"