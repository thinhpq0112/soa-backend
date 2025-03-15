package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

//type Product struct {
//	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
//	Reference  string    `json:"reference" gorm:"type:varchar(50);not null;unique"`
//	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
//	AddedDate  time.Time `json:"added_date" gorm:"type:date;default:CURRENT_DATE"`
//	Status     string    `json:"status" gorm:"type:varchar(50)"`
//	CategoryId uuid.UUID `json:"category_id" gorm:"type:uuid"`
//	Price      float64   `json:"price" gorm:"type:numeric(10,2);default:0"`
//	StockCity  string    `json:"stock_city" gorm:"type:varchar(100)"`
//	SupplierId uuid.UUID `json:"supplier_id" gorm:"type:uuid"`
//	Quantity   int       `json:"quantity" gorm:"type:int;default:0"`
//}

type Product struct {
	Id         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Reference  string    `json:"reference" gorm:"type:varchar(50);not null;unique"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	AddedDate  time.Time `json:"added_date" gorm:"type:date;default:CURRENT_DATE"`
	Status     string    `json:"status" gorm:"type:varchar(50)"`
	CategoryId uuid.UUID `json:"-" gorm:"type:uuid"`

	Price      float64   `json:"price" gorm:"type:numeric(10,2);default:0"`
	StockCity  string    `json:"stock_city" gorm:"type:varchar(100)"`
	SupplierId uuid.UUID `json:"-" gorm:"type:uuid"`

	Quantity int `json:"quantity" gorm:"type:int;default:0"`

	Category *Category `json:"category"`
	Supplier *Supplier `json:"supplier"`
}

type FilterOption struct {
	Reference  string   `json:"reference"`
	StartDate  string   `json:"start_date"`
	EndDate    string   `json:"end_date"`
	MinPrice   *float64 `json:"min_price"`
	MaxPrice   *float64 `json:"max_price"`
	Categories []string `json:"categories"`
	Suppliers  []string `json:"suppliers"`
	StockCity  []string `json:"stock"`
	Status     []string `json:"status"`
	Search     string   `json:"search"`
}

type ProductsPerCategoryResponse struct {
	CategoryName string  `json:"category_name"`
	Percentage   float64 `json:"percentage"`
}

type ProductsPerSupplierResponse struct {
	SupplierName string  `json:"supplier_name"`
	Percentage   float64 `json:"percentage"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	type Alias Product

	return json.Marshal(&struct {
		Date string `json:"added_date"`
		Alias
	}{
		Date:  p.AddedDate.Format("2006-01-02"),
		Alias: (Alias)(p),
	})
}

func (p *Product) UnmarshalJSON(data []byte) error {
	type Alias Product
	aux := &struct {
		Date string `json:"added_date"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	p.AddedDate, err = time.Parse("2006-01-02", aux.Date)
	return err
}
