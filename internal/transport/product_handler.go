package transport

import (
	"errors"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/service"
	"net/http"
	"strconv"
	"strings"
	"time"
)
import "github.com/gin-gonic/gin"

type productHandler struct {
	svc service.IProductService
}

func NewProductHandler(svc service.IProductService) *productHandler {
	return &productHandler{
		svc: svc,
	}
}

func (h *productHandler) RegisterRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/products")
	product.GET("/", h.GetProducts)
	product.GET("/:id", h.GetProductById)
	product.POST("/", h.AddProduct)
	product.PUT("/", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)

	statistics := rg.Group("/statistics")
	statistics.GET("/products-per-category", h.GetProductsPerCategory)
	statistics.GET("/products-per-supplier", h.GetProductsPerSupplier)

	product.GET("/pdf", h.GeneratePDF)
}

// @Summary Get all products
// @Description Get all products with optional filters
// @Tags products
// @Accept json
// @Produce json
// @Param page_number query int false "Page number"
// @Param limit query int false "Limit"
// @Param last_created_at query string false "Last created at"
// @Param reference query string false "Reference"
// @Param start_date query string false "Start date"
// @Param end_date query string false "End date"
// @Param min_price query float64 false "Minimum price"
// @Param max_price query float64 false "Maximum price"
// @Param categories query string false "Categories (comma-separated, e.g., Books,Electronics)"
// @Param suppliers query string false "Suppliers (comma-separated, e.g., Supplier1,Supplier2)"
// @Param stock_cities query string false "Stock cities (comma-separated, e.g., NY,LA,Chicago)"
// @Param status query string false "Status (comma-separated, e.g., Available,OutOfStock)"
// @Param search query string false "Search"
// @Success 200 {object} model.ProductListResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products [get]
func (h *productHandler) GetProducts(c *gin.Context) {
	pageNumber := parseIntQuery(c, "page_number", 1)
	limit := parseIntQuery(c, "limit", 0)
	lastCreatedAt := parseTimeQuery(c, "last_created_at")

	_, _, err := parseDateRange(c, "start_date", "end_date")
	if err != nil {
		handleBadRequest(c, err)
		return
	}

	minPrice, maxPrice, err := parsePriceRange(c, "min_price", "max_price")
	if err != nil {
		handleBadRequest(c, err)
		return
	}

	categories := parseMultiQuery(c, "categories")
	suppliers := parseMultiQuery(c, "suppliers")
	stockCities := parseMultiQuery(c, "stock_cities")
	status := parseMultiQuery(c, "status")

	options := &model.FilterOption{
		Reference:  c.Query("reference"),
		StartDate:  c.Query("start_date"),
		EndDate:    c.Query("end_date"),
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		Categories: categories,
		Suppliers:  suppliers,
		StockCity:  stockCities,
		Status:     status,
		Search:     c.Query("search"),
	}

	products, err := h.svc.GetProducts(c, pageNumber, limit, lastCreatedAt, options)
	if err != nil {
		handleErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, model.ProductListResponse{Data: products})
}

// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products/{id} [get]
func (h *productHandler) GetProductById(c *gin.Context) {
	product, err := h.svc.GetProductById(c, c.Param("id"))
	if err != nil {
		handleErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// @Summary Delete product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.ActionResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products/{id} [delete]
func (h *productHandler) DeleteProduct(c *gin.Context) {
	if err := h.svc.DeleteProduct(c, c.Param("id")); err != nil {
		handleErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// @Summary Create product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data"
// @Success 200 {object} model.ActionResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products [post]
func (h *productHandler) AddProduct(c *gin.Context) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequest(c, err)
		return
	}

	err = h.svc.AddProduct(c, product)
	if err != nil {
		handleErrorServer(c, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product added successfully"})
}

// @Summary Update product
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Product data"
// @Success 200 {object} model.ActionResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products [put]
func (h *productHandler) UpdateProduct(c *gin.Context) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		handleBadRequest(c, err)
		return
	}

	err = h.svc.UpdateProduct(c, product)
	if err != nil {
		handleErrorServer(c, err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// @Summary Get products per category
// @Description Get the number of products per category
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} model.StatPercentResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/statistics/products-per-category [get]
func (h *productHandler) GetProductsPerCategory(c *gin.Context) {
	stats, err := h.svc.GetProductsPerCategory(c)
	if err != nil {
		handleErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stats})
}

// @Summary Get products per supplier
// @Description Get the number of products per supplier
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} model.StatPercentResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /api/statistics/products-per-supplier [get]
func (h *productHandler) GetProductsPerSupplier(c *gin.Context) {
	stats, err := h.svc.GetProductsPerSupplier(c)
	if err != nil {
		handleErrorServer(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stats})
}

// @Summary Generate product report as PDF
// @Description Generates a product report in PDF format and returns it as a downloadable file
// @Tags products
// @Produce application/pdf
// @Success 200 {file} application/pdf "PDF file"
// @Failure 500 {object} model.ErrorResponse
// @Router /api/products/pdf [get]
func (h *productHandler) GeneratePDF(c *gin.Context) {
	filePath, err := h.svc.GenerateProductPDF(c)
	if err != nil {
		handleErrorServer(c, err)
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename="product_report.pdf"`)
	c.File(filePath)
}

func parseIntQuery(c *gin.Context, key string, defaultValue int) *int {
	val, err := strconv.Atoi(c.DefaultQuery(key, strconv.Itoa(defaultValue)))
	if err != nil {
		return nil
	}
	return &val
}

func parseTimeQuery(c *gin.Context, key string) *time.Time {
	val := c.Query(key)
	if val == "" {
		return nil
	}

	t, err := time.Parse("2006-01-02", val)
	if err != nil {
		return nil
	}
	return &t
}

func parseDateRange(c *gin.Context, startKey, endKey string) (*time.Time, *time.Time, error) {
	startDate := parseTimeQuery(c, startKey)
	endDate := parseTimeQuery(c, endKey)

	if startDate != nil && endDate != nil && startDate.After(*endDate) {
		return nil, nil, errors.New("invalid params: end_date must be after start_date")
	}

	return startDate, endDate, nil
}

func parsePriceRange(c *gin.Context, minKey, maxKey string) (*float64, *float64, error) {
	minPrice := parseFloatQuery(c, minKey)
	maxPrice := parseFloatQuery(c, maxKey)

	if minPrice != nil && maxPrice != nil && *minPrice > *maxPrice {
		return nil, nil, errors.New("invalid params: min_price must be less than or equal to max_price")
	}

	return minPrice, maxPrice, nil
}

func parseFloatQuery(c *gin.Context, key string) *float64 {
	val := c.Query(key)
	if val == "" {
		return nil
	}

	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil
	}
	return &f
}

func parseMultiQuery(c *gin.Context, key string) []string {
	vals := c.QueryArray(key)
	if len(vals) == 1 && strings.Contains(vals[0], ",") {
		return strings.Split(vals[0], ",")
	}
	return vals
}

func handleBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func handleErrorServer(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
