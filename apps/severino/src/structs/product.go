package structs

import (
	"time"
)

type Product struct {
	ID                string                   `json:"id"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at,omitempty"`
	StartedSellingAt  time.Time                `json:"started_selling_at"`
	FinishedSellingAt time.Time                `json:"finished_selling_at"`
	UsageExpiresAt    time.Time                `json:"usage_expires_at"`
	Name              string                   `json:"name" validate:"required"`
	Description       string                   `json:"description"`
	Stock             int32                    `json:"stock"`
	SKU               uint64                   `json:"sku"`
	Image             string                   `json:"image"`
	IsPublished       bool                     `json:"is_published"`
	Company           string                   `json:"company" validate:"required"`
	ProductType       string                   `json:"product_type" validate:"required"`
	Items             []map[string]interface{} `json:"items" validate:"required"`
	PaymentsTypes     []PaymentType            `json:"payments_types" validate:"required"`
}

type PaymentType struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int32  `json:"price"`
	Installments int32  `json:"installments"`
}
