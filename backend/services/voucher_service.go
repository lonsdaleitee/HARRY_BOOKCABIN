package services

import (
	"database/sql"
	"fmt"

	"airline-voucher-backend/models"
	"airline-voucher-backend/utils"
)

// VoucherService handles voucher-related business logic
type VoucherService struct {
	db models.Database
}

// NewVoucherService creates a new VoucherService instance
func NewVoucherService(db models.Database) *VoucherService {
	return &VoucherService{
		db: db,
	}
}

// CheckVoucherExists checks if a voucher already exists for the given flight and date
func (s *VoucherService) CheckVoucherExists(flightNumber, date string) (bool, error) {
	query := `SELECT COUNT(*) FROM vouchers WHERE flight_number = ? AND flight_date = ?`

	var count int
	err := s.db.QueryRow(query, flightNumber, date).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check voucher existence: %w", err)
	}

	return count > 0, nil
}

// GenerateVoucher generates a new voucher with 3 random seats
func (s *VoucherService) GenerateVoucher(req *models.GenerateVoucherRequest) (*models.GenerateVoucherResponse, error) {
	// Validate aircraft type
	if !utils.ValidateAircraftType(req.Aircraft) {
		return nil, fmt.Errorf("invalid aircraft type: %s", req.Aircraft)
	}

	// Validate date format
	if !utils.ValidateDateFormat(req.Date) {
		return nil, fmt.Errorf("invalid date format: %s (expected YYYY-MM-DD)", req.Date)
	}

	// Check if voucher already exists
	exists, err := s.CheckVoucherExists(req.FlightNumber, req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to check voucher existence: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("voucher already exists for flight %s on %s", req.FlightNumber, req.Date)
	}

	// Generate random seats
	seats, err := utils.GenerateRandomSeats(req.Aircraft)
	if err != nil {
		return nil, fmt.Errorf("failed to generate seats: %w", err)
	}

	// Save voucher to database
	err = s.saveVoucher(req, seats)
	if err != nil {
		return nil, fmt.Errorf("failed to save voucher: %w", err)
	}

	return &models.GenerateVoucherResponse{
		Success: true,
		Seats:   seats,
	}, nil
}

// saveVoucher saves the voucher to the database
func (s *VoucherService) saveVoucher(req *models.GenerateVoucherRequest, seats []string) error {
	query := `
		INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	currentTime := models.GetCurrentTimestamp()

	_, err := s.db.Exec(
		query,
		req.Name,
		req.ID,
		req.FlightNumber,
		req.Date,
		req.Aircraft,
		seats[0],
		seats[1],
		seats[2],
		currentTime,
	)

	return err
}

// GetVoucherByFlightAndDate retrieves a voucher by flight number and date
func (s *VoucherService) GetVoucherByFlightAndDate(flightNumber, date string) (*models.Voucher, error) {
	query := `
		SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at
		FROM vouchers WHERE flight_number = ? AND flight_date = ?
	`

	voucher := &models.Voucher{}
	err := s.db.QueryRow(query, flightNumber, date).Scan(
		&voucher.ID,
		&voucher.CrewName,
		&voucher.CrewID,
		&voucher.FlightNumber,
		&voucher.FlightDate,
		&voucher.AircraftType,
		&voucher.Seat1,
		&voucher.Seat2,
		&voucher.Seat3,
		&voucher.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get voucher: %w", err)
	}

	return voucher, nil
}
