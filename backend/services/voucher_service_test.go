package services

import (
	"testing"

	"airline-voucher-backend/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVoucherService(t *testing.T) {
	// Simple test to verify service creation
	service := NewVoucherService(nil)

	require.NotNil(t, service)
}

func TestVoucherService_ValidationLogic(t *testing.T) {
	// Test validation logic without invoking database operations
	tests := []struct {
		name        string
		request     *models.GenerateVoucherRequest
		expectError bool
		errorMsg    string
	}{
		{
			name: "Invalid aircraft type",
			request: &models.GenerateVoucherRequest{
				Name:         "John Doe",
				ID:           "12345",
				FlightNumber: "GA102",
				Date:         "2025-07-12",
				Aircraft:     "Invalid Aircraft",
			},
			expectError: true,
			errorMsg:    "invalid aircraft type",
		},
		{
			name: "Invalid date format",
			request: &models.GenerateVoucherRequest{
				Name:         "John Doe",
				ID:           "12345",
				FlightNumber: "GA102",
				Date:         "12-07-2025", // Wrong format
				Aircraft:     "ATR",
			},
			expectError: true,
			errorMsg:    "invalid date format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewVoucherService(nil)
			_, err := service.GenerateVoucher(tt.request)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
