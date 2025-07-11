# 🐳 Docker Deployment Guide

This guide provides comprehensive instructions for running the Airline Voucher Seat Assignment System using Docker.

> **📚 Documentation Navigation:**
> - **Quick Start**: See [DOCKER_SUMMARY.md](./DOCKER_SUMMARY.md) for fast commands
> - **Project Overview**: See [README.md](./README.md) for general information
> - **This Guide**: Detailed deployment, troubleshooting, and maintenance

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Production Deployment](#production-deployment)
- [Development Environment](#development-environment)
- [Docker Architecture](#docker-architecture)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [Maintenance](#maintenance)

## 🔧 Prerequisites

Before running the application with Docker, ensure you have:

- **Docker Engine** 20.10.0+ installed
- **Docker Compose** V2+ installed
- At least **2GB RAM** available
- Ports **3000** and **8080** available

### Installation Links
- [Docker Desktop for Mac](https://docs.docker.com/docker-for-mac/install/)
- [Docker Desktop for Windows](https://docs.docker.com/docker-for-windows/install/)
- [Docker Engine for Linux](https://docs.docker.com/engine/install/)

## 🚀 Quick Start

### Option 1: Using Convenience Scripts (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd HARRY_BOOKCABIN

# Start production environment
./docker-start.sh

# Or start development environment
./docker-start-dev.sh
```

### Option 2: Manual Docker Compose

```bash
# Production deployment
docker-compose up --build -d

# Development deployment
docker-compose -f docker-compose.dev.yml up --build -d
```

### ✅ Verify Installation

After starting the containers, verify the services are running:

```bash
# Check container status
docker-compose ps

# Test backend health
curl http://localhost:8080/health

# Test frontend
curl http://localhost:3000/health
```

## 🏭 Production Deployment

### Production Architecture

The production setup includes:
- **Backend**: Go application with optimized binary
- **Frontend**: React app served by Nginx
- **Database**: SQLite with persistent volume
- **Networking**: Internal Docker network with health checks

### Starting Production Environment

```bash
# Using convenience script
./docker-start.sh

# Or manually
docker-compose up --build -d
```

### Production Environment Variables

Create a `.env` file in the root directory:

```env
# Backend Configuration
GIN_MODE=release
DB_PATH=/root/data/vouchers.db

# Frontend Configuration
VITE_API_URL=http://localhost:8080/api
```

### Production Services

| Service | URL | Description |
|---------|-----|-------------|
| Frontend | http://localhost:3000 | React application |
| Backend API | http://localhost:8080/api | Go REST API |
| Health Check | http://localhost:8080/health | Backend health status |

## 🛠️ Development Environment

### Development Features

- **Hot Reloading**: Frontend and backend auto-reload on code changes
- **Debug Mode**: Go Gin framework in debug mode
- **Volume Mounting**: Source code mounted for live editing
- **Development Dependencies**: Includes dev tools and debuggers

### Starting Development Environment

```bash
# Using convenience script
./docker-start-dev.sh

# Or manually
docker-compose -f docker-compose.dev.yml up --build -d
```

### Development Services

| Service | URL | Description |
|---------|-----|-------------|
| Frontend Dev | http://localhost:3000 | Vite dev server with HMR |
| Backend Dev | http://localhost:8080/api | Go app with auto-reload |

### Development Workflow

1. **Make code changes** in your editor
2. **Frontend**: Changes auto-reload via Vite HMR
3. **Backend**: Container restarts automatically on Go file changes
4. **Database**: Data persists in Docker volume

## 🏗️ Docker Architecture

### Container Structure

```
┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │
│   (React+Nginx) │    │   (Go+SQLite)   │
│   Port: 3000    │    │   Port: 8080    │
└─────────────────┘    └─────────────────┘
         │                       │
         └───────────────────────┘
                     │
            ┌─────────────────┐
            │ Docker Network  │
            │ (airline-network)│
            └─────────────────┘
```

### Volume Management

- **backend_data**: Persists SQLite database
- **Source mounting** (dev only): Live code editing

### Health Checks

Both containers include health checks:
- **Interval**: 30 seconds
- **Timeout**: 10 seconds
- **Retries**: 3 attempts
- **Start Period**: 10 seconds

## ⚙️ Configuration

### Environment Variables

#### Backend Configuration
```env
GIN_MODE=release|debug
DB_PATH=/path/to/database.db
PORT=8080
```

#### Frontend Configuration
```env
VITE_API_URL=http://localhost:8080/api
```

### Port Configuration

Default ports can be changed in `docker-compose.yml`:

```yaml
services:
  frontend:
    ports:
      - "3000:80"  # Change 3000 to desired port
  backend:
    ports:
      - "8080:8080"  # Change 8080 to desired port
```

### Database Configuration

SQLite database is stored in Docker volumes:
- **Production**: `/root/data/vouchers.db`
- **Development**: `/app/data/vouchers.db`

## 🔍 Troubleshooting

### Common Issues

#### Port Already in Use
```bash
# Find process using port
lsof -i :3000
lsof -i :8080

# Kill process or change ports in docker-compose.yml
```

#### Container Won't Start
```bash
# Check container logs
docker-compose logs backend
docker-compose logs frontend

# Check container status
docker-compose ps
```

#### Database Issues
```bash
# Check database volume
docker volume ls
docker volume inspect harry_bookcabin_backend_data

# Reset database (⚠️ Data loss)
docker-compose down -v
```

#### Build Issues
```bash
# Clean build
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

### Debug Commands

```bash
# Execute shell in running container
docker-compose exec backend sh
docker-compose exec frontend sh

# View real-time logs
docker-compose logs -f

# Check resource usage
docker stats

# Inspect network
docker network ls
docker network inspect harry_bookcabin_airline-network
```

### Performance Issues

```bash
# Check resource usage
docker stats

# Allocate more memory to Docker Desktop
# Docker Desktop > Settings > Resources > Advanced
```

## 🔧 Maintenance

### Regular Maintenance Tasks

#### Update Dependencies
```bash
# Stop containers
./docker-stop.sh

# Rebuild with latest dependencies
docker-compose build --no-cache
docker-compose up -d
```

#### Backup Database
```bash
# Create database backup
docker-compose exec backend cp /root/data/vouchers.db /root/data/vouchers_backup_$(date +%Y%m%d).db

# Copy backup to host
docker cp airline-voucher-backend:/root/data/vouchers_backup_$(date +%Y%m%d).db ./backups/
```

#### Clean Up Resources
```bash
# Using convenience script
./docker-clean.sh

# Or manually
docker system prune -f
docker volume prune -f
```

### Monitoring

#### Container Health
```bash
# Check health status
docker-compose ps

# Health check endpoints
curl http://localhost:8080/health
curl http://localhost:3000/health
```

#### Logs Management
```bash
# View logs
docker-compose logs

# Follow logs in real-time
docker-compose logs -f

# View logs for specific service
docker-compose logs backend
docker-compose logs frontend
```

## 📝 Available Scripts

### Production Scripts
- `./docker-start.sh` - Start production environment
- `./docker-stop.sh` - Stop all containers
- `./docker-clean.sh` - Clean up Docker resources

### Development Scripts
- `./docker-start-dev.sh` - Start development environment

### Manual Commands
```bash
# Production
docker-compose up -d              # Start detached
docker-compose down              # Stop and remove containers
docker-compose logs -f           # View logs
docker-compose ps               # Check status

# Development
docker-compose -f docker-compose.dev.yml up -d
docker-compose -f docker-compose.dev.yml down
docker-compose -f docker-compose.dev.yml logs -f
```

## 🛡️ Security Considerations

### Production Security
- Containers run as non-root user where possible
- Security headers configured in Nginx
- Database stored in persistent volume
- Internal network isolation

### Development Security
- Development containers should not be used in production
- Debug mode exposes additional information
- Source code mounted as volumes

## 📊 Performance Optimization

### Production Optimizations
- Multi-stage Docker builds for smaller images
- Nginx gzip compression enabled
- Static asset caching
- Health checks for automatic recovery

### Resource Requirements
- **Minimum**: 2GB RAM, 1 CPU core
- **Recommended**: 4GB RAM, 2 CPU cores
- **Storage**: 10GB available disk space

## 🔄 CI/CD Integration

### Docker Build Pipeline
```yaml
# Example GitHub Actions
- name: Build and test
  run: |
    docker-compose build
    docker-compose up -d
    # Run tests
    docker-compose down
```

## 📞 Support

For issues related to Docker deployment:

1. Check the [Troubleshooting](#troubleshooting) section
2. Review container logs: `docker-compose logs`
3. Verify system requirements
4. Check Docker and Docker Compose versions

---

## ✅ Deployment Checklist

- [ ] Docker and Docker Compose installed
- [ ] Ports 3000 and 8080 available
- [ ] Sufficient system resources
- [ ] Scripts have execute permissions
- [ ] Environment variables configured
- [ ] Health checks passing
- [ ] Database persistence working
- [ ] Frontend and backend communication verified

🎉 **Your Airline Voucher System is now ready to run with Docker!**
