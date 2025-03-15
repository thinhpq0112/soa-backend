package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/thinhpq0112/soa-backend/internal/model"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) GetProductById(ctx context.Context, id string) (model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductRepo) AddProduct(ctx context.Context, product model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepo) GetProducts(ctx context.Context, pageNumber, limit *int, lastCreatedAt *time.Time, options *model.FilterOption) ([]model.Product, error) {
	args := m.Called(ctx, pageNumber, limit, lastCreatedAt, options)
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepo) DeleteProduct(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductRepo) UpdateProduct(ctx context.Context, product model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepo) GetProductsPerCategory(ctx context.Context) ([]model.ProductsPerCategoryResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.ProductsPerCategoryResponse), args.Error(1)
}

func (m *MockProductRepo) GetProductsPerSupplier(ctx context.Context) ([]model.ProductsPerSupplierResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.ProductsPerSupplierResponse), args.Error(1)
}
