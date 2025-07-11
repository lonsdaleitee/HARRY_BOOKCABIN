#!/bin/bash

# Build and start development containers
echo "🚀 Building and starting Airline Voucher System in Development Mode..."

# Stop any running containers
echo "📦 Stopping existing containers..."
docker-compose -f docker-compose.dev.yml down

# Build and start containers
echo "🔨 Building and starting development containers..."
docker-compose -f docker-compose.dev.yml up --build -d

# Show container status
echo "📊 Container status:"
docker-compose -f docker-compose.dev.yml ps

# Show logs
echo "📝 Showing logs (Ctrl+C to exit):"
docker-compose -f docker-compose.dev.yml logs -f
