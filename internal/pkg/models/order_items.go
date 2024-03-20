package models

import "time"

// OrderItem model struct
type OrderItems struct {
	ID        int        `gorm:"column:id"`
	OrderID   int        `gorm:"column:order_id"`
	GroceryID int64      `gorm:"column:grocery_id"`
	Quantity  int        `gorm:"column:quantity"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

// TableName returns the table name for the order_items model
func (*OrderItems) TableName() string {
	return "order_items"
}
