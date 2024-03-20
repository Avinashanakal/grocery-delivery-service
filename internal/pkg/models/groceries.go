package models

import "time"

// Grocery model struct
type Grocery struct {
	ID          int        `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	Price       float64    `gorm:"column:price"`
	Quantity    int        `gorm:"column:quantity"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

// TableName returns the table name for the groceries model
func (Grocery) TableName() string {
	return "groceries"
}
