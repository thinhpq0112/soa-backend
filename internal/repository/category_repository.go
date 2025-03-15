package repository

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryById(ctx context.Context, id string) (model.Category, error)
	AddCategory(ctx context.Context, category model.Category) error
	UpdateCategory(ctx context.Context, category model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepo) GetCategoryById(ctx context.Context, id string) (model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *CategoryRepo) AddCategory(ctx context.Context, category model.Category) error {
	return r.db.WithContext(ctx).Create(&category).Error
}

func (r *CategoryRepo) UpdateCategory(ctx context.Context, category model.Category) error {
	return r.db.WithContext(ctx).Save(&category).Error
}

func (r *CategoryRepo) DeleteCategory(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Category{}, id).Error
}
