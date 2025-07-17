package handlers

import (
	"fmt"
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

// GetVoucher handles POST /api/voucher requests to get existing voucher
func (h *VoucherHandler) GetVoucher(c *gin.Context) {
	var req models.GetVoucherRequest

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

	voucher, err := h.service.GetVoucher(req.FlightNumber, req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to get voucher",
			Message: err.Error(),
		})
		return
	}

	response := models.GetVoucherResponse{
		Voucher: voucher,
		Exists:  voucher != nil,
	}

	c.JSON(http.StatusOK, response)
}

// RegenerateSeat handles POST /api/regenerate-seat requests
func (h *VoucherHandler) RegenerateSeat(c *gin.Context) {
	var req models.RegenerateSeatRequest

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
			Message: "FlightNumber, date, and seatPosition are required",
		})
		return
	}

	// Validate seat position
	if req.SeatPosition < 1 || req.SeatPosition > 3 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid seat position",
			Message: "Seat position must be 1, 2, or 3",
		})
		return
	}

	response, err := h.service.RegenerateSeat(&req)
	if err != nil {
		// Check if it's a business logic error (voucher not found)
		if err.Error() == "no voucher found for flight "+req.FlightNumber+" on "+req.Date {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Error:   "Voucher not found",
				Message: err.Error(),
			})
			return
		}

		// Check if it's a validation error
		if err.Error() == fmt.Sprintf("invalid seat position: %d (must be 1, 2, or 3)", req.SeatPosition) {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Error:   "Invalid seat position",
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to regenerate seat",
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
