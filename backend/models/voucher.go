package models

import (
	"database/sql"
	"time"
)

// Voucher represents a voucher assignment in the database
type Voucher struct {
	ID           int    `json:"id" db:"id"`
	CrewName     string `json:"crew_name" db:"crew_name"`
	CrewID       string `json:"crew_id" db:"crew_id"`
	FlightNumber string `json:"flight_number" db:"flight_number"`
	FlightDate   string `json:"flight_date" db:"flight_date"`
	AircraftType string `json:"aircraft_type" db:"aircraft_type"`
	Seat1        string `json:"seat1" db:"seat1"`
	Seat2        string `json:"seat2" db:"seat2"`
	Seat3        string `json:"seat3" db:"seat3"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

// CheckVoucherRequest represents the request to check if vouchers exist
type CheckVoucherRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

// CheckVoucherResponse represents the response for checking vouchers
type CheckVoucherResponse struct {
	Exists bool `json:"exists"`
}

// GenerateVoucherRequest represents the request to generate vouchers
type GenerateVoucherRequest struct {
	Name         string `json:"name" binding:"required"`
	ID           string `json:"id" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Aircraft     string `json:"aircraft" binding:"required"`
}

// GenerateVoucherResponse represents the response for generating vouchers
type GenerateVoucherResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// Database interface for testing
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Close() error
}

// GetCurrentTimestamp returns the current timestamp in ISO format
func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
