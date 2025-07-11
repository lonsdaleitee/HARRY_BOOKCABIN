package handlers

import (
	"net/http"

	"airline-voucher-backend/models"
	"airline-voucher-backend/services"

	"github.com/gin-gonic/gin"
)

// VoucherHandler handles voucher-related HTTP requests
type VoucherHandler struct {
	service *services.VoucherService
}

// NewVoucherHandler creates a new VoucherHandler instance
func NewVoucherHandler(service *services.VoucherService) *VoucherHandler {
	return &VoucherHandler{
		service: service,
	}
}

// CheckVoucher handles POST /api/check requests
func (h *VoucherHandler) CheckVoucher(c *gin.Context) {
	var req models.CheckVoucherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	// Validate required fields
	if req.FlightNumber == "" || req.Date == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Missing required fields",
			Message: "Both flightNumber and date are required",
		})
		return
	}

	exists, err := h.service.CheckVoucherExists(req.FlightNumber, req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to check voucher",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.CheckVoucherResponse{
		Exists: exists,
	})
}

// GenerateVoucher handles POST /api/generate requests
func (h *VoucherHandler) GenerateVoucher(c *gin.Context) {
	var req models.GenerateVoucherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request body",
			Message: err.Error(),
		})
		return
	}

	// Validate required fields
	if req.Name == "" || req.ID == "" || req.FlightNumber == "" || req.Date == "" || req.Aircraft == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Missing required fields",
			Message: "All fields (name, id, flightNumber, date, aircraft) are required",
		})
		return
	}

	response, err := h.service.GenerateVoucher(&req)
	if err != nil {
		// Check if it's a business logic error (voucher already exists)
		if err.Error() == "voucher already exists for flight "+req.FlightNumber+" on "+req.Date {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error:   "Voucher already exists",
				Message: err.Error(),
			})
			return
		}

		// Check if it's a validation error
		if err.Error() == "invalid aircraft type: "+req.Aircraft {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "Invalid aircraft type",
				Message: err.Error(),
			})
			return
		}

		if err.Error() == "invalid date format: "+req.Date+" (expected YYYY-MM-DD)" {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "Invalid date format",
				Message: err.Error(),
			})
			return
		}

		// Internal server error
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to generate voucher",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// HealthCheck handles GET /health requests
func (h *VoucherHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Airline voucher service is running",
	})
}
