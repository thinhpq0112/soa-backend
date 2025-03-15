package repository

import (
	"context"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"gorm.io/gorm"
)

type ISupplierRepo interface {
	GetSuppliers(ctx context.Context) ([]model.Supplier, error)
	GetSupplierById(ctx context.Context, id string) (model.Supplier, error)
	AddSupplier(ctx context.Context, supplier model.Supplier) error
	UpdateSupplier(ctx context.Context, supplier model.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error
}

type supplierRepo struct {
	db *gorm.DB
}

func NewSupplierRepo(db *gorm.DB) *supplierRepo {
	return &supplierRepo{db: db}
}

func (r *supplierRepo) GetSuppliers(ctx context.Context) ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := r.db.WithContext(ctx).Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepo) GetSupplierById(ctx context.Context, id string) (model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&supplier).Error
	return supplier, err
}

func (r *supplierRepo) AddSupplier(ctx context.Context, supplier model.Supplier) error {
	return r.db.WithContext(ctx).Create(&supplier).Error
}

func (r *supplierRepo) UpdateSupplier(ctx context.Context, supplier model.Supplier) error {
	return r.db.WithContext(ctx).Save(&supplier).Error
}

func (r *supplierRepo) DeleteSupplier(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Supplier{}, id).Error
}
