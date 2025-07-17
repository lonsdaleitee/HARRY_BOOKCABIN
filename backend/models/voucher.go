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
	Message string `json:"message"`
}

// GetVoucherRequest represents the request to get existing vouchers
type GetVoucherRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

// GetVoucherResponse represents the response for getting vouchers
type GetVoucherResponse struct {
	Voucher *Voucher `json:"voucher"`
	Exists  bool     `json:"exists"`
}

// RegenerateSeatRequest represents the request to regenerate a single seat
type RegenerateSeatRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
	SeatPosition int    `json:"seatPosition" binding:"required,min=1,max=3"` // 1, 2, or 3
}

// RegenerateSeatResponse represents the response for regenerating a single seat
type RegenerateSeatResponse struct {
	Success  bool     `json:"success"`
	NewSeat  string   `json:"newSeat"`
	AllSeats []string `json:"allSeats"` // All three seats after regeneration
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
