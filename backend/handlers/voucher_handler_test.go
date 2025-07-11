package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"airline-voucher-backend/models"
	"airline-voucher-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRouter(handler *VoucherHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/check", handler.CheckVoucher)
		api.POST("/generate", handler.GenerateVoucher)
	}

	router.GET("/health", handler.HealthCheck)

	return router
}

func TestVoucherHandler_CheckVoucher_Validation(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
	}{
		{
			name: "Invalid request body",
			requestBody: map[string]interface{}{
				"invalidField": "value",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing required fields",
			requestBody: models.CheckVoucherRequest{
				FlightNumber: "GA102",
				// Date is missing
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceWrapper := &services.VoucherService{}
			handler := &VoucherHandler{service: serviceWrapper}

			router := setupTestRouter(handler)

			// Prepare request
			jsonBody, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/check", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestVoucherHandler_GenerateVoucher_Validation(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
	}{
		{
			name: "Missing required fields",
			requestBody: models.GenerateVoucherRequest{
				Name:         "John Doe",
				ID:           "12345",
				FlightNumber: "GA102",
				// Date and Aircraft are missing
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceWrapper := &services.VoucherService{}
			handler := &VoucherHandler{service: serviceWrapper}

			router := setupTestRouter(handler)

			// Prepare request
			var jsonBody []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				jsonBody = []byte(str)
			} else {
				jsonBody, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req, err := http.NewRequest("POST", "/api/generate", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestVoucherHandler_HealthCheck(t *testing.T) {
	serviceWrapper := &services.VoucherService{}
	handler := &VoucherHandler{service: serviceWrapper}

	router := setupTestRouter(handler)

	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, "healthy", response["status"])
	assert.Contains(t, response["message"], "running")
}
