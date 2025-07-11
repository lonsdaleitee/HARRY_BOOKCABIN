package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAircraftConfig(t *testing.T) {
	tests := []struct {
		name          string
		aircraftType  string
		expectedRows  int
		expectedSeats []string
		expectError   bool
	}{
		{
			name:          "ATR aircraft",
			aircraftType:  "ATR",
			expectedRows:  18,
			expectedSeats: []string{"A", "C", "D", "F"},
			expectError:   false,
		},
		{
			name:          "Airbus 320 aircraft",
			aircraftType:  "Airbus 320",
			expectedRows:  32,
			expectedSeats: []string{"A", "B", "C", "D", "E", "F"},
			expectError:   false,
		},
		{
			name:          "Boeing 737 Max aircraft",
			aircraftType:  "Boeing 737 Max",
			expectedRows:  32,
			expectedSeats: []string{"A", "B", "C", "D", "E", "F"},
			expectError:   false,
		},
		{
			name:         "Unknown aircraft type",
			aircraftType: "Unknown",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := GetAircraftConfig(tt.aircraftType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, config)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, config)
			assert.Equal(t, tt.expectedRows, config.Rows)
			assert.Equal(t, tt.expectedSeats, config.Seats)
		})
	}
}

func TestGenerateRandomSeats(t *testing.T) {
	tests := []struct {
		name         string
		aircraftType string
		expectError  bool
	}{
		{
			name:         "ATR aircraft",
			aircraftType: "ATR",
			expectError:  false,
		},
		{
			name:         "Airbus 320 aircraft",
			aircraftType: "Airbus 320",
			expectError:  false,
		},
		{
			name:         "Boeing 737 Max aircraft",
			aircraftType: "Boeing 737 Max",
			expectError:  false,
		},
		{
			name:         "Unknown aircraft type",
			aircraftType: "Unknown",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seats, err := GenerateRandomSeats(tt.aircraftType)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, seats)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, seats)
			assert.Len(t, seats, 3)

			// Check that all seats are unique
			uniqueSeats := make(map[string]bool)
			for _, seat := range seats {
				assert.False(t, uniqueSeats[seat], "Seat %s should be unique", seat)
				uniqueSeats[seat] = true
			}

			// Check that all seats are valid for the aircraft type
			config, _ := GetAircraftConfig(tt.aircraftType)
			for _, seat := range seats {
				assert.Regexp(t, `^\d+[A-F]$`, seat, "Seat %s should match the pattern", seat)

				// Extract row and seat letter
				var row int
				var letter string

				// Parse row number and seat letter - handle both single and double digit rows
				if len(seat) >= 2 {
					if len(seat) == 2 {
						// Single digit row (e.g., "1A")
						row = int(seat[0] - '0')
						letter = string(seat[1])
					} else if len(seat) == 3 {
						// Double digit row (e.g., "10A")
						row = int(seat[0]-'0')*10 + int(seat[1]-'0')
						letter = string(seat[2])
					}

					assert.True(t, row >= 1 && row <= config.Rows, "Row %d should be within valid range 1-%d", row, config.Rows)

					// Check if seat letter is valid for this aircraft
					validLetter := false
					for _, validSeat := range config.Seats {
						if letter == validSeat {
							validLetter = true
							break
						}
					}
					assert.True(t, validLetter, "Seat letter %s should be valid for %s", letter, tt.aircraftType)
				}
			}
		})
	}
}

func TestValidateAircraftType(t *testing.T) {
	tests := []struct {
		name         string
		aircraftType string
		expected     bool
	}{
		{
			name:         "Valid ATR",
			aircraftType: "ATR",
			expected:     true,
		},
		{
			name:         "Valid Airbus 320",
			aircraftType: "Airbus 320",
			expected:     true,
		},
		{
			name:         "Valid Boeing 737 Max",
			aircraftType: "Boeing 737 Max",
			expected:     true,
		},
		{
			name:         "Invalid aircraft type",
			aircraftType: "Invalid",
			expected:     false,
		},
		{
			name:         "Empty string",
			aircraftType: "",
			expected:     false,
		},
		{
			name:         "Case sensitive test",
			aircraftType: "atr",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateAircraftType(tt.aircraftType)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateDateFormat(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected bool
	}{
		{
			name:     "Valid date format",
			date:     "2025-07-12",
			expected: true,
		},
		{
			name:     "Valid date format with leading zeros",
			date:     "2025-01-01",
			expected: true,
		},
		{
			name:     "Invalid date format - DD-MM-YYYY",
			date:     "12-07-2025",
			expected: false,
		},
		{
			name:     "Invalid date format - MM/DD/YYYY",
			date:     "07/12/2025",
			expected: false,
		},
		{
			name:     "Invalid date format - missing year",
			date:     "07-12",
			expected: false,
		},
		{
			name:     "Invalid date format - empty string",
			date:     "",
			expected: false,
		},
		{
			name:     "Invalid date format - invalid month",
			date:     "2025-13-01",
			expected: false,
		},
		{
			name:     "Invalid date format - invalid day",
			date:     "2025-02-30",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateDateFormat(tt.date)
			assert.Equal(t, tt.expected, result)
		})
	}
}
