# Airline Voucher Backend

A Go backend service for the airline voucher seat assignment system using Gin Gonic framework and SQLite database.

## Features

- **REST API**: Clean RESTful endpoints for voucher management
- **SQLite Database**: Lightweight, file-based database for voucher storage
- **CORS Support**: Configured for frontend integration
- **Input Validation**: Comprehensive request validation
- **Error Handling**: Structured error responses
- **Unit Tests**: Comprehensive test coverage
- **Modular Architecture**: Clean separation of concerns

## Tech Stack

- **Go 1.21** - Programming language
- **Gin Gonic** - Web framework
- **SQLite3** - Database
- **Testify** - Testing framework

## Project Structure

```
backend/
├── config/           # Configuration and database setup
├── handlers/         # HTTP request handlers
├── models/          # Data models and structures
├── services/        # Business logic layer
├── utils/           # Utility functions (seat generation, etc.)
├── main.go          # Application entry point
└── go.mod           # Go module dependencies
```

## API Endpoints

### Health Check
- **GET** `/health` - Service health check

### Voucher Endpoints
- **POST** `/api/check` - Check if vouchers exist for a flight/date
- **POST** `/api/generate` - Generate new voucher assignments

## Database Schema

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

## Aircraft Seat Layouts

- **ATR**: 18 rows, seats A,C,D,F (72 total seats)
- **Airbus 320**: 32 rows, seats A,B,C,D,E,F (192 total seats)
- **Boeing 737 Max**: 32 rows, seats A,B,C,D,E,F (192 total seats)

## Getting Started

### Prerequisites

- Go 1.21 or higher
- SQLite3 (included with Go sqlite3 driver)

### Installation

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Download dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

### Available Commands

- `go run main.go` - Start the development server
- `go build` - Build the application binary
- `go test ./...` - Run all tests
- `go test -v ./...` - Run tests with verbose output
- `go test -cover ./...` - Run tests with coverage report

## API Usage Examples

### Check if vouchers exist
```bash
curl -X POST http://localhost:8080/api/check \
  -H "Content-Type: application/json" \
  -d '{
    "flightNumber": "GA102",
    "date": "2025-07-12"
  }'
```

Response:
```json
{
  "exists": false
}
```

### Generate vouchers
```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Sarah",
    "id": "98123",
    "flightNumber": "GA102",
    "date": "2025-07-12",
    "aircraft": "Airbus 320"
  }'
```

Response:
```json
{
  "success": true,
  "seats": ["3B", "7C", "14D"]
}
```

## Error Handling

The API returns structured error responses:

```json
{
  "error": "Error type",
  "message": "Detailed error description"
}
```

### Common HTTP Status Codes

- `200` - Success
- `400` - Bad Request (validation errors)
- `409` - Conflict (voucher already exists)
- `500` - Internal Server Error

## Security Features

- **Parameterized SQL Queries**: Protection against SQL injection
- **Input Validation**: Comprehensive request validation
- **CORS Configuration**: Secure cross-origin resource sharing

## Configuration

The application uses the following default configuration:

- **Port**: 8080
- **Database**: `./vouchers.db`
- **CORS Origin**: `http://localhost:3000` (frontend)

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./utils
go test ./services
go test ./handlers
```

### Test Coverage

The test suite covers:
- Seat generation algorithms
- Aircraft configuration validation
- Date format validation
- Service layer business logic
- HTTP handler request/response processing
- Error handling scenarios

## Development

### Adding New Aircraft Types

1. Update the aircraft configuration in `utils/seats.go`
2. Add validation in `utils/seats.go`
3. Update tests in `utils/seats_test.go`

### Database Migrations

For schema changes:
1. Update the schema in `config/config.go`
2. Handle existing data migration if needed
3. Update model structures in `models/voucher.go`

## Deployment

### Local Development
```bash
go run main.go
```

### Production Build
```bash
go build -o airline-voucher-backend main.go
./airline-voucher-backend
```

### Docker (Optional)
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

## Performance Considerations

- SQLite database with indexed queries for fast lookups
- Efficient random seat generation algorithm
- Minimal memory footprint
- Stateless design for horizontal scaling

## Monitoring

The service provides:
- Health check endpoint for load balancer integration
- Structured logging for debugging
- Error tracking and reporting

## Contributing

1. Follow Go conventions and best practices
2. Add tests for new features
3. Update documentation for API changes
4. Use parameterized queries for database operations
