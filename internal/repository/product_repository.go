package repository

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"gorm.io/gorm"
	"time"
)

type IProductRepo interface {
	GetProducts(ctx context.Context, pageNumber, limit *int, lastCreatedAt *time.Time, options *model.FilterOption) ([]model.Product, error)
	GetProductById(ctx context.Context, id string) (model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	UpdateProduct(ctx context.Context, product model.Product) error

	AddProduct(ctx context.Context, product model.Product) error
	GetProductsPerCategory(ctx context.Context) ([]model.ProductsPerCategoryResponse, error)
	GetProductsPerSupplier(ctx context.Context) ([]model.ProductsPerSupplierResponse, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db: db}
}

const maxPageLimit = 100
const defaultSizeLimit = 20

func (p *productRepo) GetProducts(ctx context.Context, pageNumber *int, limit *int, lastCreatedAt *time.Time, options *model.FilterOption) ([]model.Product, error) {
	var products []model.Product

	query := p.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier")

	if len(options.Categories) > 0 || options.Search != "" {
		query = query.Joins("JOIN categories ON categories.id = products.category_id")
	}

	if len(options.Suppliers) > 0 || options.Search != "" {
		query = query.Joins("JOIN suppliers ON suppliers.id = products.supplier_id")
	}

	if len(options.Categories) > 0 {
		query = query.Where("categories.name IN (?)", options.Categories)
	}

	if len(options.Suppliers) > 0 {
		query = query.Where("suppliers.name IN (?)", options.Suppliers)
	}

	if options.Reference != "" {
		query = query.Where("reference = ?", options.Reference)
	}

	if options.StartDate != "" {
		query = query.Where("added_date >= ?", options.StartDate)
	}

	if options.EndDate != "" {
		query = query.Where("added_date <= ?", options.EndDate)
	}

	if len(options.Status) > 0 {
		query = query.Where("status IN (?)", options.Status)
	}

	if len(options.StockCity) > 0 {
		query = query.Where("stock_city IN (?)", options.StockCity)
	}

	if options.MinPrice != nil {
		query = query.Where("price >= ?", *options.MinPrice)
	}

	if options.MaxPrice != nil {
		query = query.Where("price <= ?", *options.MaxPrice)
	}

	if options.Search != "" {
		search := "%" + options.Search + "%"
		query = query.Where(`
			(reference ILIKE ? 
			OR stock_city ILIKE ? 
			OR categories.name ILIKE ? 
			OR suppliers.name ILIKE ? 
			OR products.name ILIKE ? 
			OR status ILIKE ? 
			OR CAST(price AS TEXT) ILIKE ?)`,
			search, search, search, search, search, search, search,
		)
	}

	if limit == nil || *limit <= 0 {
		d := defaultSizeLimit
		limit = &d
	}
	if *limit > maxPageLimit {
		*limit = maxPageLimit
	}

	if lastCreatedAt != nil {
		query = query.Where("added_date > ?", *lastCreatedAt)
	} else if pageNumber != nil {
		offset := (*pageNumber - 1) * *limit
		query = query.Offset(offset)
	}

	query = query.Limit(*limit)

	err := query.Find(&products).Error
	return products, err
}

func (p *productRepo) GetProductById(ctx context.Context, id string) (model.Product, error) {
	var product model.Product
	err := p.db.WithContext(ctx).
		Preload("Category").
		Preload("Supplier").
		Where("id = ?", id).
		First(&product).Error
	return product, err
}

func (p *productRepo) UpdateProduct(ctx context.Context, product model.Product) error {
	return p.db.WithContext(ctx).Updates(&product).Error
}

func (p *productRepo) DeleteProduct(ctx context.Context, id string) error {
	return p.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Product{}).Error
}

func (p *productRepo) AddProduct(ctx context.Context, product model.Product) error {
	return p.db.Create(&product).Error
}

func (p *productRepo) GetProductsPerCategory(ctx context.Context) ([]model.ProductsPerCategoryResponse, error) {
	var results []model.ProductsPerCategoryResponse
	err := p.db.WithContext(ctx).
		Table("products").
		Select("categories.name as category_name, COUNT(*) * 100.0 / SUM(COUNT(*)) OVER() as percentage").
		Joins("left join categories on products.category_id = categories.id").
		Group("categories.name").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (p *productRepo) GetProductsPerSupplier(ctx context.Context) ([]model.ProductsPerSupplierResponse, error) {
	var results []model.ProductsPerSupplierResponse
	err := p.db.WithContext(ctx).
		Table("products").
		Select("suppliers.name as supplier_name, COUNT(*) * 100.0 / SUM(COUNT(*)) OVER() as percentage").
		Joins("left join suppliers on products.supplier_id = suppliers.id").
		Group("suppliers.name").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
