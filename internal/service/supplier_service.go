package service

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"github.com/thinhpq0112/soa-backend/internal/repository"
)

type ISupplierService interface {
	GetSuppliers(ctx context.Context) ([]model.Supplier, error)
	GetSupplierById(ctx context.Context, id string) (model.Supplier, error)
	AddSupplier(ctx context.Context, supplier model.Supplier) error
	UpdateSupplier(ctx context.Context, supplier model.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
}

type supplierService struct {
	repo repository.ISupplierRepo
}

func NewSupplierService(repo repository.ISupplierRepo) *supplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetSuppliers(ctx context.Context) ([]model.Supplier, error) {
	return s.repo.GetSuppliers(ctx)
}

func (s *supplierService) GetSupplierById(ctx context.Context, id string) (model.Supplier, error) {
	return s.repo.GetSupplierById(ctx, id)
}

func (s *supplierService) AddSupplier(ctx context.Context, supplier model.Supplier) error {
	return s.repo.AddSupplier(ctx, supplier)
}

func (s *supplierService) UpdateSupplier(ctx context.Context, supplier model.Supplier) error {
	return s.repo.UpdateSupplier(ctx, supplier)
}

func (s *supplierService) DeleteSupplier(ctx context.Context, id string) error {
	return s.repo.DeleteSupplier(ctx, id)
}
