package models

import (
	"encoding/json"
	"time"
)

// OrderStatus type for ENUM
type OrderStatus string

const (
	Placed    OrderStatus = "PLACED"
	Shipped   OrderStatus = "SHIPPED"
	Delivered OrderStatus = "DELIVERED"
	Cancelled OrderStatus = "CANCELLED"
)

// Order model struct
type Order struct {
	ID           int         `gorm:"column:id" json:"id"`
	CustomerID   int64       `gorm:"column:customer_id" json:"customer_id"`
	OrderDate    time.Time   `gorm:"column:order_date" json:"order_date"`
	DeliveryDate time.Time   `gorm:"column:delivery_date" json:"delivery_date"`
	TotalPrice   float64     `gorm:"column:total_price" json:"total_price"`
	Status       OrderStatus `gorm:"column:status" json:"status"`
	CreatedAt    time.Time   `gorm:"column:created_at" json:"-"`
	UpdatedAt    *time.Time  `gorm:"column:updated_at" json:"-"`
}

// TableName returns the table name for the orders model
func (Order) TableName() string {
	return "orders"
}

func (o *Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	temp := struct {
		OrderDate    string `json:"order_date"`
		DeliveryDate string `json:"delivery_date"`
		CreatedAt    string `json:"created_at"`
		*Alias
	}{
		OrderDate:    o.OrderDate.Format("2006-01-02"),
		DeliveryDate: o.DeliveryDate.Format("2006-01-02T15:04:05"),
		CreatedAt:    o.CreatedAt.Format("2006-01-02T15:04:05"),
		Alias:        (*Alias)(o),
	}

	return json.Marshal(temp)
}
