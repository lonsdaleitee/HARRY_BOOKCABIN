#!/bin/bash

# Airline Voucher Backend Build and Run Script

set -e

echo "🚀 Airline Voucher Backend Setup"
echo "================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✅ Go version: $GO_VERSION"

# Navigate to backend directory
cd "$(dirname "$0")"

# Download dependencies
echo "📦 Downloading dependencies..."
go mod download

# Run tests
echo "🧪 Running tests..."
go test ./... -v

# Build the application
echo "🔨 Building application..."
go build -o airline-voucher-backend main.go

echo ""
echo "✅ Build completed successfully!"
echo ""
echo "🎯 To start the server:"
echo "   ./airline-voucher-backend"
echo ""
echo "📝 Or run in development mode:"
echo "   go run main.go"
echo ""
echo "🌐 Server will be available at:"
echo "   http://localhost:8080"
echo ""
echo "🏥 Health check:"
echo "   curl http://localhost:8080/health"
echo ""
