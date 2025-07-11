package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// AircraftConfig represents the configuration for an aircraft type
type AircraftConfig struct {
	Rows  int
	Seats []string
}

// GetAircraftConfig returns the seat configuration for a given aircraft type
func GetAircraftConfig(aircraftType string) (*AircraftConfig, error) {
	configs := map[string]AircraftConfig{
		"ATR": {
			Rows:  18,
			Seats: []string{"A", "C", "D", "F"},
		},
		"Airbus 320": {
			Rows:  32,
			Seats: []string{"A", "B", "C", "D", "E", "F"},
		},
		"Boeing 737 Max": {
			Rows:  32,
			Seats: []string{"A", "B", "C", "D", "E", "F"},
		},
	}

	config, exists := configs[aircraftType]
	if !exists {
		return nil, fmt.Errorf("unknown aircraft type: %s", aircraftType)
	}

	return &config, nil
}

// GenerateRandomSeats generates 3 unique random seats for the given aircraft type
func GenerateRandomSeats(aircraftType string) ([]string, error) {
	config, err := GetAircraftConfig(aircraftType)
	if err != nil {
		return nil, err
	}

	// Generate all possible seats
	var allSeats []string
	for row := 1; row <= config.Rows; row++ {
		for _, seat := range config.Seats {
			allSeats = append(allSeats, fmt.Sprintf("%d%s", row, seat))
		}
	}

	// Use current time as seed for randomness
	rand.Seed(time.Now().UnixNano())

	// Shuffle the seats and take the first 3
	shuffled := make([]string, len(allSeats))
	copy(shuffled, allSeats)

	// Fisher-Yates shuffle
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	// Return the first 3 seats
	return shuffled[:3], nil
}

// ValidateAircraftType checks if the aircraft type is valid
func ValidateAircraftType(aircraftType string) bool {
	validTypes := []string{"ATR", "Airbus 320", "Boeing 737 Max"}
	for _, validType := range validTypes {
		if aircraftType == validType {
			return true
		}
	}
	return false
}

// ValidateDateFormat validates the date format (YYYY-MM-DD)
func ValidateDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
