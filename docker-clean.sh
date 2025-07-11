#!/bin/bash

# Clean up Docker resources
echo "🧹 Cleaning up Docker resources..."

# Stop all containers
echo "📦 Stopping all containers..."
docker-compose down
docker-compose -f docker-compose.dev.yml down

# Remove containers, networks, images, and volumes
echo "🗑️  Removing containers, networks, and images..."
docker-compose down --rmi all --volumes --remove-orphans
docker-compose -f docker-compose.dev.yml down --rmi all --volumes --remove-orphans

# Remove unused Docker resources
echo "🧽 Removing unused Docker resources..."
docker system prune -f

echo "✅ Docker cleanup completed!"
