#!/bin/bash

# Build and start production containers
echo "🚀 Building and starting Airline Voucher System with Docker..."

# Stop any running containers
echo "📦 Stopping existing containers..."
docker-compose down

# Build and start containers
echo "🔨 Building and starting containers..."
docker-compose up --build -d

# Show container status
echo "📊 Container status:"
docker-compose ps

# Show logs
echo "📝 Showing logs (Ctrl+C to exit):"
docker-compose logs -f
