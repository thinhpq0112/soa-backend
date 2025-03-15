package service

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/repository"
)

type ICategoryService interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryById(ctx context.Context, id string) (model.Category, error)
	AddCategory(ctx context.Context, category model.Category) error
	UpdateCategory(ctx context.Context, category model.Category) error
	DeleteCategory(ctx context.Context, id string) error
}

type CategoryService struct {
	repo repository.ICategoryRepo
}

func NewCategoryService(repo repository.ICategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetCategories(ctx)
}

func (s *CategoryService) GetCategoryById(ctx context.Context, id string) (model.Category, error) {
	return s.repo.GetCategoryById(ctx, id)
}

func (s *CategoryService) AddCategory(ctx context.Context, category model.Category) error {
	return s.repo.AddCategory(ctx, category)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, category model.Category) error {
	return s.repo.UpdateCategory(ctx, category)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	return s.repo.DeleteCategory(ctx, id)
}
