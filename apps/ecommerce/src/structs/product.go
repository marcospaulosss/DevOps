package structs

import (
	"fmt"
	"time"

	"backend/libs/json"
)

type Product struct {
	ID                string    `json:"id" db:"id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	StartedSellingAt  time.Time `json:"started_selling_at" db:"started_selling_at"`
	FinishedSellingAt time.Time `json:"finished_selling_at" db:"finished_selling_at"`
	UsageExpiresAt    time.Time `json:"usage_expires_at" db:"usage_expires_at"`
	Name              string    `json:"name" db:"name"`
	Description       string    `json:"description" db:"description"`
	Stock             int32     `json:"stock" db:"stock"`
	SKU               uint64    `json:"sku" db:"sku"`
	Image             string    `json:"image" db:"image"`
	IsPublished       bool      `json:"is_published" db:"is_published"`
	Company           string    `json:"company" db:"company"`
	ProductType       string    `json:"product_type" db:"product_type"`
	Items             *string   `json:"items" db:"items"`
	PaymentsTypes     *string   `json:"payments_types" db:"payments_types"`
	History           *string   `json:"history" db:"history"`
}

type PaymentType struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Price        int32  `json:"price" db:"price"`
	Installments int32  `json:"installments" db:"installments"`
}

func (this Product) GetHistory() ([]Product, error) {
	var productHistory []Product
	if this.History == nil {
		return productHistory, fmt.Errorf("O produto nao tem history e por isso GetHistory() vai retornar nil.")
	}
	history := *this.History
	err := json.Unmarshal([]byte(history), &productHistory)
	return productHistory, err
}

func (this Product) GetPaymentsTypes() ([]PaymentType, error) {
	var paymentsTypes []PaymentType
	if this.PaymentsTypes == nil {
		return paymentsTypes, fmt.Errorf("O produto nao tem tipos de pagamento e por isso GetPaymentsTypes() vai retornar nil.")
	}
	types := *this.PaymentsTypes
	err := json.Unmarshal([]byte(types), &paymentsTypes)
	return paymentsTypes, err
}
