package service

import (
	"codeberg.org/go-pdf/fpdf"
	"context"
	"fmt"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/repository"

	"time"
)

type IProductService interface {
	GetProducts(ctx context.Context, pageNumber, limit *int, lastCreatedAt *time.Time, option *model.FilterOption) ([]model.Product, error)
	GetProductById(ctx context.Context, id string) (model.Product, error)
	AddProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, id string) error

	GetProductsPerCategory(ctx context.Context) ([]model.ProductsPerCategoryResponse, error)
	GetProductsPerSupplier(ctx context.Context) ([]model.ProductsPerSupplierResponse, error)
	GenerateProductPDF(ctx context.Context) (string, error)
}

type productService struct {
	repo repository.IProductRepo
}

func NewProductService(repo repository.IProductRepo) *productService {
	return &productService{repo: repo}
}

func (s *productService) GetProducts(ctx context.Context, pageNumber, limit *int, lastCreatedAt *time.Time, option *model.FilterOption) ([]model.Product, error) {
	return s.repo.GetProducts(ctx, pageNumber, limit, lastCreatedAt, option)
}

func (s *productService) GetProductById(ctx context.Context, id string) (model.Product, error) {
	return s.repo.GetProductById(ctx, id)
}

func (s *productService) AddProduct(ctx context.Context, product model.Product) error {
	return s.repo.AddProduct(ctx, product)
}

func (s *productService) UpdateProduct(ctx context.Context, product model.Product) error {
	return s.repo.UpdateProduct(ctx, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.DeleteProduct(ctx, id)
}

func (s *productService) GetProductsPerCategory(ctx context.Context) ([]model.ProductsPerCategoryResponse, error) {
	return s.repo.GetProductsPerCategory(ctx)
}

func (s *productService) GetProductsPerSupplier(ctx context.Context) ([]model.ProductsPerSupplierResponse, error) {
	return s.repo.GetProductsPerSupplier(ctx)
}

func (s *productService) GenerateProductPDF(ctx context.Context) (string, error) {
	products, err := s.repo.GetProducts(ctx, nil, nil, nil, &model.FilterOption{})
	if err != nil {
		return "", err
	}

	pdf := fpdf.New("L", "mm", "A3", "")
	pdf.SetMargins(10, 10, 10)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 18)
	pdf.CellFormat(0, 12, "Product Report", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	headers := []string{
		"Product Reference", "Product Name", "Date Added", "Status",
		"Product Category", "Price (EUR)", "Stock Location (City)",
		"Supplier", "Availability Quantity",
	}

	colWidths := []float64{40, 55, 40, 30, 45, 35, 40, 45, 40}

	pdf.SetFont("Arial", "B", 10)
	for i, h := range headers {
		pdf.CellFormat(colWidths[i], 8, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 9)
	for _, product := range products {
		pdf.CellFormat(40, 8, product.Reference, "1", 0, "L", false, 0, "")
		pdf.CellFormat(55, 8, product.Name, "1", 0, "L", false, 0, "")

		dateStr := "N/A"
		if !product.AddedDate.IsZero() {
			dateStr = product.AddedDate.Format("2006-01-02")
		}
		pdf.CellFormat(40, 8, dateStr, "1", 0, "C", false, 0, "")

		pdf.CellFormat(30, 8, product.Status, "1", 0, "C", false, 0, "")

		categoryName := "Unknown"
		if product.Category != nil {
			categoryName = product.Category.Name
		}
		pdf.CellFormat(45, 8, categoryName, "1", 0, "L", false, 0, "")

		pdf.CellFormat(35, 8, fmt.Sprintf("%.2f", product.Price), "1", 0, "C", false, 0, "")

		pdf.CellFormat(40, 8, product.StockCity, "1", 0, "C", false, 0, "")

		supplierName := "Unknown"
		if product.Supplier != nil {
			supplierName = product.Supplier.Name
		}
		pdf.CellFormat(45, 8, supplierName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 8, fmt.Sprintf("%d", product.Quantity), "1", 0, "C", false, 0, "")

		pdf.Ln(-1)
	}

	filePath := "product_report.pdf"
	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
