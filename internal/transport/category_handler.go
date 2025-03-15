package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/service"
	"net/http"
)

type CategoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) RegisterRoutes(rg *gin.RouterGroup) {
	category := rg.Group("/categories")
	category.GET("/", h.GetCategories)
	category.GET("/:id", h.GetCategoryById)
	category.POST("/", h.AddCategory)
	category.PUT("/:id", h.UpdateCategory)
	category.DELETE("/:id", h.DeleteCategory)
}

// @Summary Get all categories
// @Description Retrieve a list of all categories
// @Tags categories
// @Produce json
// @Success 200 {array} model.Category
// @Failure 500 {object} model.ErrorResponse
// @Router /api/categories [get]
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// @Summary Get a category by ID
// @Description Retrieve a category by its unique ID
// @Tags categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/categories/{id} [get]
func (h *CategoryHandler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	category, err := h.service.GetCategoryById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// @Summary Add a new category
// @Description Create a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body model.Category true "Category data"
// @Success 201 {object} model.Category
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/categories [post]
func (h *CategoryHandler) AddCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AddCategory(c.Request.Context(), category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

// @Summary Update a category
// @Description Update an existing category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body model.Category true "Updated category data"
// @Success 200 {object} model.Category
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.Id = uuid.MustParse(id)
	if err := h.service.UpdateCategory(c.Request.Context(), category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// @Summary Delete a category
// @Description Delete a category by ID
// @Tags categories
// @Param id path string true "Category ID"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
