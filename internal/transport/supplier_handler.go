package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/service"
	"net/http"
)

type SupplierHandler struct {
	service service.ISupplierService
}

func NewSupplierHandler(service service.ISupplierService) *SupplierHandler {
	return &SupplierHandler{service: service}
}

func (h *SupplierHandler) RegisterRoutes(rg *gin.RouterGroup) {
	supplier := rg.Group("/suppliers")
	supplier.GET("/", h.GetSuppliers)
	supplier.GET("/:id", h.GetSupplierById)
	supplier.POST("/", h.AddSupplier)
	supplier.PUT("/:id", h.UpdateSupplier)
	supplier.DELETE("/:id", h.DeleteSupplier)
}

// @Summary Get all suppliers
// @Description Retrieve a list of all suppliers
// @Tags suppliers
// @Produce json
// @Success 200 {array} model.Supplier
// @Failure 500 {object} model.ErrorResponse
// @Router /api/suppliers [get]
func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	suppliers, err := h.service.GetSuppliers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

// @Summary Get a supplier by ID
// @Description Retrieve a supplier by its unique ID
// @Tags suppliers
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} model.Supplier
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/suppliers/{id} [get]
func (h *SupplierHandler) GetSupplierById(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.service.GetSupplierById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// @Summary Add a new supplier
// @Description Create a new supplier
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body model.Supplier true "Supplier data"
// @Success 201 {object} model.Supplier
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/suppliers [post]
func (h *SupplierHandler) AddSupplier(c *gin.Context) {
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddSupplier(c.Request.Context(), supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, supplier)
}

// @Summary Update a supplier
// @Description Update an existing supplier by ID
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Param supplier body model.Supplier true "Updated supplier data"
// @Success 200 {object} model.Supplier
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/suppliers/{id} [put]
func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier model.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	supplier.Id = uuid.MustParse(id)
	if err := h.service.UpdateSupplier(c.Request.Context(), supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

// @Summary Delete a supplier
// @Description Delete a supplier by ID
// @Tags suppliers
// @Param id path string true "Supplier ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/suppliers/{id} [delete]
func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteSupplier(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
