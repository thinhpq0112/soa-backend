package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/thinhpq0112/soa-backend/internal/repository/mocks"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thinhpq0112/soa-backend/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestMain(m *testing.M) {
	viper.SetConfigFile("../../.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}
	os.Exit(m.Run())
}

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	return gormDB, mock
}

// Test query
func TestAddProduct(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewProductRepo(db)

	mockUUID := uuid.New()
	product := model.Product{
		Reference:  "Test Reference",
		Name:       "Test Product",
		Status:     "Test Status",
		CategoryId: uuid.Must(uuid.Parse("94d0da61-0bbe-4be8-8435-2b72f03a29ea")),
		Price:      100.0,
		StockCity:  "Test Stock",
		SupplierId: uuid.Must(uuid.Parse("4f8ce93f-46c2-4d20-8a27-92fdbf6ee464")),
		Quantity:   9,
		AddedDate:  time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "products" \("reference","name","status","category_id","price","stock_city","supplier_id","quantity","added_date"\) 
		VALUES \(\$1,\$2,\$3,\$4,\$5,\$6,\$7,\$8,\$9\) RETURNING "id","added_date"`).
		WithArgs(product.Reference, product.Name, product.Status, product.CategoryId, product.Price, product.StockCity, product.SupplierId, product.Quantity, product.AddedDate).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockUUID))
	mock.ExpectCommit()

	err := repo.AddProduct(context.Background(), product)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetProductByID(t *testing.T) {
	mockRepo := new(mocks.MockProductRepo)

	productID := uuid.New()
	categoryID := uuid.New()
	supplierID := uuid.New()

	expectedProduct := model.Product{
		Id:         productID,
		Reference:  "Test Reference",
		Name:       "Test Product",
		AddedDate:  time.Now(),
		Status:     "Available",
		CategoryId: categoryID,
		Price:      100.0,
		StockCity:  "Test Stock",
		SupplierId: supplierID,
		Quantity:   50,
		Category: &model.Category{
			Id:   categoryID,
			Name: "Electronics",
		},
		Supplier: &model.Supplier{
			Id:   supplierID,
			Name: "Tech Supplier",
		},
	}

	mockRepo.On("GetProductById", mock.Anything, expectedProduct.Id.String()).
		Return(expectedProduct, nil)

	product, err := mockRepo.GetProductById(context.Background(), expectedProduct.Id.String())

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepo)

	productID := uuid.New()
	categoryID := uuid.New()
	supplierID := uuid.New()

	updatedProduct := model.Product{
		Id:         productID,
		Reference:  "Updated Reference",
		Name:       "Updated Product",
		AddedDate:  time.Now(),
		Status:     "Out of Stock",
		CategoryId: categoryID,
		Price:      120.0,
		StockCity:  "Updated Location",
		SupplierId: supplierID,
		Quantity:   30,
	}

	mockRepo.On("UpdateProduct", mock.Anything, updatedProduct).
		Return(nil)

	err := mockRepo.UpdateProduct(context.Background(), updatedProduct)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(mocks.MockProductRepo)

	productID := uuid.New()

	mockRepo.On("DeleteProduct", mock.Anything, productID.String()).
		Return(nil)

	err := mockRepo.DeleteProduct(context.Background(), productID.String())

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
