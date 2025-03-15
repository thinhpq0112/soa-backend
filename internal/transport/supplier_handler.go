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

func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	suppliers, err := h.service.GetSuppliers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suppliers)
}

func (h *SupplierHandler) GetSupplierById(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.service.GetSupplierById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supplier)
}

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

func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteSupplier(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
