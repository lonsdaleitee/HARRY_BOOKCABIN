#!/bin/bash

# Stop all containers and clean up
echo "🛑 Stopping Airline Voucher System containers..."

# Stop production containers
echo "📦 Stopping production containers..."
docker-compose down

# Stop development containers
echo "📦 Stopping development containers..."
docker-compose -f docker-compose.dev.yml down

echo "✅ All containers stopped successfully!"

# Optional: Remove images (uncomment if needed)
# echo "🗑️  Removing Docker images..."
# docker-compose down --rmi all
# docker-compose -f docker-compose.dev.yml down --rmi all
