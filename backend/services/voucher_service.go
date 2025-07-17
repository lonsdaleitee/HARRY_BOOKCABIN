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

// GetVoucher retrieves an existing voucher for the given flight and date
func (s *VoucherService) GetVoucher(flightNumber, date string) (*models.Voucher, error) {
	query := `SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at 
			  FROM vouchers WHERE flight_number = ? AND flight_date = ? LIMIT 1`

	var voucher models.Voucher
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
			return nil, nil // Voucher not found
		}
		return nil, fmt.Errorf("failed to get voucher: %w", err)
	}

	return &voucher, nil
}

// RegenerateSeat regenerates a single seat for an existing voucher
func (s *VoucherService) RegenerateSeat(req *models.RegenerateSeatRequest) (*models.RegenerateSeatResponse, error) {
	// Validate seat position
	if req.SeatPosition < 1 || req.SeatPosition > 3 {
		return nil, fmt.Errorf("invalid seat position: %d (must be 1, 2, or 3)", req.SeatPosition)
	}

	// Get existing voucher
	voucher, err := s.GetVoucher(req.FlightNumber, req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to get voucher: %w", err)
	}

	if voucher == nil {
		return nil, fmt.Errorf("no voucher found for flight %s on %s", req.FlightNumber, req.Date)
	}

	// Get current seats
	currentSeats := []string{voucher.Seat1, voucher.Seat2, voucher.Seat3}

	// Generate all possible seats for this aircraft
	allPossibleSeats, err := utils.GetAllSeats(voucher.AircraftType)
	if err != nil {
		return nil, fmt.Errorf("failed to get available seats: %w", err)
	}

	// Filter out currently assigned seats (except the one we're regenerating)
	var availableSeats []string

	for _, seat := range allPossibleSeats {
		isOccupied := false
		for i, currentSeat := range currentSeats {
			if seat == currentSeat && i != (req.SeatPosition-1) {
				isOccupied = true
				break
			}
		}
		if !isOccupied {
			availableSeats = append(availableSeats, seat)
		}
	}

	if len(availableSeats) == 0 {
		return nil, fmt.Errorf("no available seats to regenerate")
	}

	// Generate a new random seat from available options
	newSeat, err := utils.GenerateRandomSeat(availableSeats)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new seat: %w", err)
	}

	// Update the specific seat in the database
	var updateQuery string
	var seatColumn string

	switch req.SeatPosition {
	case 1:
		seatColumn = "seat1"
	case 2:
		seatColumn = "seat2"
	case 3:
		seatColumn = "seat3"
	}

	updateQuery = fmt.Sprintf("UPDATE vouchers SET %s = ? WHERE flight_number = ? AND flight_date = ?", seatColumn)

	_, err = s.db.Exec(updateQuery, newSeat, req.FlightNumber, req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to update seat: %w", err)
	}

	// Update current seats array with new seat
	currentSeats[req.SeatPosition-1] = newSeat

	return &models.RegenerateSeatResponse{
		Success:  true,
		NewSeat:  newSeat,
		AllSeats: currentSeats,
	}, nil
}
