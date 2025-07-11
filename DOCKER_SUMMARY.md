# 🐳 Docker Quick Reference

## 🚀 One-Command Setup

```bash
# Production (Recommended for demos/testing)
./docker-start.sh

# Development (For coding with hot reload)  
./docker-start-dev.sh

# Stop everything
./docker-stop.sh
```

## 🌐 Access Points

| Service | URL | Description |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | React app (Nginx in prod, Vite in dev) |
| **Backend API** | http://localhost:8080/api | Go REST API |
| **Health Check** | http://localhost:8080/health | Service status |

## 🔧 Management Commands

```bash
# View logs
docker-compose logs -f

# Check status  
docker-compose ps

# Restart single service
docker-compose restart backend
docker-compose restart frontend

# Clean rebuild
./docker-stop.sh && ./docker-start.sh

# Cleanup resources
./docker-clean.sh
```

## 🏗️ What's Included

✅ **Production Environment**
- Multi-stage optimized builds
- Nginx serving React static files
- Go binary with SQLite persistence
- Health checks and auto-recovery

✅ **Development Environment**  
- Hot reloading for both frontend/backend
- Source code volume mounting
- Debug logging enabled
- Real-time file watching

✅ **DevOps Features**
- Persistent database volumes
- Internal Docker networking
- Environment-specific configs
- Container health monitoring

## 📁 Key Files

```
├── docker-compose.yml         # Production config
├── docker-compose.dev.yml     # Development config  
├── docker-start.sh           # Production start
├── docker-start-dev.sh       # Development start
├── docker-stop.sh            # Stop all
├── docker-clean.sh           # Cleanup
├── backend/Dockerfile        # Go backend image
├── frontend/Dockerfile       # React + Nginx image
└── frontend/nginx.conf       # Production web server
```

## 🆘 Quick Troubleshooting

```bash
# Port conflicts
lsof -i :3000 :8080

# Container issues  
docker-compose logs backend
docker-compose logs frontend

# Database reset (⚠️ data loss)
docker-compose down -v

# Network issues
docker network ls
docker network prune
```

## 📖 Need More Help?

- **Detailed Setup**: [DOCKER.md](./DOCKER.md)
- **Project Overview**: [README.md](./README.md)
- **Requirements**: See main README for dependencies
