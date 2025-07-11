# 🛫 Airline Voucher Seat Assignment System

A full-stack web application for managing airline voucher seat assignments with random seat generation based on aircraft configurations.

## 📋 Project Overview

This application fulfills all requirements from the airline voucher seat assignment specification:

- **Frontend**: React + TypeScript with modern UI components
- **Backend**: Go with Gin Gonic framework and SQLite database
- **Features**: Random seat generation, duplicate prevention, comprehensive validation

## 🏗️ Architecture

```
HARRY_BOOKCABIN/
├── frontend/           # React TypeScript application
│   ├── src/
│   │   ├── api/       # API service layer
│   │   ├── components/ # React components
│   │   ├── store/     # State management (Jotai)
│   │   ├── types/     # TypeScript definitions
│   │   ├── utils/     # Utility functions
│   │   └── test/      # Test files
│   └── package.json
├── backend/           # Go backend application
│   ├── config/       # Database configuration
│   ├── handlers/     # HTTP request handlers
│   ├── models/       # Data models
│   ├── services/     # Business logic
│   ├── utils/        # Utility functions
│   ├── main.go       # Application entry point
│   └── go.mod        # Go dependencies
└── BOOKCABIN_REQUIREMENT.MD
```

## ✨ Features Implemented

### ✅ Frontend Features
- **Modern React UI** with TypeScript and Vite
- **Form Validation** using Zod schemas
- **State Management** with Jotai atoms
- **Responsive Design** with custom CSS
- **Error Handling** with user-friendly messages
- **Date Utilities** for format conversion
- **Comprehensive Testing** with Vitest and React Testing Library

### ✅ Backend Features
- **Gin Gonic Web Framework** for high-performance HTTP handling
- **SQLite Database** with indexed queries for performance
- **Parameterized SQL Queries** for security against SQL injection
- **CORS Support** for frontend integration
- **Comprehensive Validation** for all input data
- **Structured Error Responses** with appropriate HTTP status codes
- **Modular Architecture** with clean separation of concerns
- **Unit Tests** covering all major components
- **Integration Tests** for end-to-end functionality

### ✅ Aircraft Configurations
- **ATR**: 18 rows, seats A,C,D,F (72 total seats)
- **Airbus 320**: 32 rows, seats A,B,C,D,E,F (192 total seats)
- **Boeing 737 Max**: 32 rows, seats A,B,C,D,E,F (192 total seats)

### ✅ API Endpoints

#### POST /api/check
Check if vouchers already exist for a flight/date combination.

**Request:**
```json
{
  "flightNumber": "GA102",
  "date": "2025-07-12"
}
```

**Response:**
```json
{
  "exists": true
}
```

#### POST /api/generate
Generate 3 random seats for a flight and save to database.

**Request:**
```json
{
  "name": "Sarah",
  "id": "98123",
  "flightNumber": "GA102",
  "date": "2025-07-12",
  "aircraft": "Airbus 320"
}
```

**Response:**
```json
{
  "success": true,
  "seats": ["3B", "7C", "14D"]
}
```

#### GET /health
Health check endpoint for monitoring.

**Response:**
```json
{
  "status": "healthy",
  "message": "Airline voucher service is running"
}
```

## 🚀 Getting Started

### Prerequisites
- **Node.js** 18+ for frontend
- **Go** 1.21+ for backend
- **SQLite3** (included with Go driver)

### 🐳 Docker Setup (Recommended)

The application is fully containerized with optimized production and development environments:

#### 🚀 Quick Start

```bash
# Production Environment (Optimized builds, Nginx serving)
./docker-start.sh

# Development Environment (Hot reloading, debug mode)
./docker-start-dev.sh

# Stop all containers
./docker-stop.sh

# Clean up resources
./docker-clean.sh
```

#### 🌐 Service Access
| Environment | Frontend | Backend API | Health Check |
|-------------|----------|-------------|--------------|
| Production  | [localhost:3000](http://localhost:3000) | [localhost:8080/api](http://localhost:8080/api) | [localhost:8080/health](http://localhost:8080/health) |
| Development | [localhost:3000](http://localhost:3000) | [localhost:8080/api](http://localhost:8080/api) | [localhost:8080/health](http://localhost:8080/health) |

#### 🏗️ Docker Features
- **Multi-stage builds** for production optimization
- **Health checks** with automatic recovery
- **Persistent volumes** for database storage  
- **Hot reloading** in development
- **Environment-specific configurations**
- **Container networking** for secure communication

#### 📚 Documentation Structure
- **[DOCKER_SUMMARY.md](./DOCKER_SUMMARY.md)** - Quick commands & reference
- **[DOCKER.md](./DOCKER.md)** - Comprehensive deployment guide

### 🎨 Frontend Setup (Manual)

```bash
cd frontend
npm install
npm run dev
```

The frontend will be available at `http://localhost:3000`

### ⚙️ Backend Setup (Manual)

```bash
cd backend
go mod download
go run main.go
```

The backend will be available at `http://localhost:8080`

### 🔧 Development Scripts

**Frontend:**
```bash
npm run dev        # Start development server
npm run build      # Build for production
npm run test       # Run tests
npm run test:watch # Run tests in watch mode
```

**Backend:**
```bash
go run main.go     # Start development server
go build          # Build binary
go test ./...     # Run all tests
./build.sh        # Build and test script
./test_integration.sh # Integration tests
```

## 🧪 Testing

### Frontend Testing
- **Unit Tests**: Component logic and utilities
- **Integration Tests**: User interactions and API integration
- **Type Safety**: Full TypeScript coverage
- **Coverage**: Comprehensive test coverage with Vitest

### Backend Testing
- **Unit Tests**: Business logic and utilities
- **Integration Tests**: Full API workflow testing
- **Database Tests**: SQLite operations and constraints
- **Error Handling**: Comprehensive error scenario testing

**Run all tests:**
```bash
# Frontend
cd frontend && npm test

# Backend
cd backend && go test ./... -v

# Integration tests
cd backend && ./test_integration.sh
```

## 🛠️ Database Schema

```sql
CREATE TABLE vouchers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    crew_name TEXT NOT NULL,
    crew_id TEXT NOT NULL,
    flight_number TEXT NOT NULL,
    flight_date TEXT NOT NULL,
    aircraft_type TEXT NOT NULL,
    seat1 TEXT NOT NULL,
    seat2 TEXT NOT NULL,
    seat3 TEXT NOT NULL,
    created_at TEXT NOT NULL
);

CREATE INDEX idx_flight_date ON vouchers(flight_number, flight_date);
```

## 📡 API Documentation

### Error Responses
All endpoints return structured error responses:

```json
{
  "error": "Error Type",
  "message": "Detailed error description"
}
```

### HTTP Status Codes
- `200` - Success
- `400` - Bad Request (validation errors)
- `409` - Conflict (duplicate voucher)
- `500` - Internal Server Error

## 🔒 Security Features

- **Parameterized SQL Queries** prevent SQL injection
- **Input Validation** on all endpoints
- **CORS Configuration** for secure frontend integration
- **Error Message Sanitization** to prevent information leakage

## 🚀 Deployment

### 🐳 Docker Deployment (Recommended)

**Production:**
```bash
# Start production environment
./docker-start.sh

# Services available at:
# Frontend: http://localhost:3000
# Backend: http://localhost:8080/api
```

**Development:**
```bash
# Start development environment with hot reloading
./docker-start-dev.sh
```

**Management:**
```bash
# Stop all containers
./docker-stop.sh

# Clean up resources
./docker-clean.sh
```

📖 **See [DOCKER.md](./DOCKER.md) for comprehensive Docker documentation**

### Manual Deployment

**Development:**
```bash
# Terminal 1: Start backend
cd backend && go run main.go

# Terminal 2: Start frontend
cd frontend && npm run dev
```

**Production:**
```bash
# Build frontend
cd frontend && npm run build

# Build backend
cd backend && go build -o airline-voucher-backend main.go

# Deploy static files and binary
```

## 📊 Performance Features

- **Indexed Database Queries** for fast lookups
- **Efficient Random Generation** algorithm
- **Minimal Memory Footprint** in Go backend
- **Optimized Frontend Bundling** with Vite

## 🔍 Monitoring

- **Health Check Endpoint** (`/health`)
- **Structured Logging** for debugging
- **Error Tracking** and reporting
- **Database Query Optimization**

## 🤝 Contributing

1. Follow Go and React best practices
2. Add tests for new features
3. Update documentation for API changes
4. Use parameterized queries for database operations

## 📝 License

This project is built for the airline voucher seat assignment requirements and includes all specified functionality with additional enhancements for production readiness.

---

## ✅ Requirements Checklist

### Frontend Requirements ✅
- [x] Crew name and ID input fields
- [x] Flight number and date input (DD-MM-YY format)
- [x] Aircraft type dropdown (ATR, Airbus 320, Boeing 737 Max)
- [x] Generate Vouchers button
- [x] API integration for check and generate endpoints
- [x] Display 3 randomly chosen seats
- [x] Error handling for duplicate vouchers

### Backend Requirements ✅
- [x] POST /api/check endpoint
- [x] POST /api/generate endpoint
- [x] SQLite database with vouchers table
- [x] Parameterized SQL queries
- [x] Aircraft seat layout configurations
- [x] 3 unique random seat generation
- [x] Duplicate prevention logic
- [x] Modular handler architecture

### Bonus Features ✅
- [x] Parameterized SQL for injection prevention
- [x] Clean error message formatting
- [x] Modular backend handlers for maintainability
- [x] Unit tests for all components
- [x] Integration tests for full workflow
- [x] Comprehensive documentation
- [x] Production-ready code structure

🎉 **All requirements successfully implemented!**
