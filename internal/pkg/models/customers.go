package models

import "time"

// Customer model struct
type Customers struct {
	ID          int        `gorm:"column:id"`
	FirstName   string     `gorm:"column:first_name"`
	LastName    string     `gorm:"column:last_name"`
	Email       string     `gorm:"column:email"`
	PhoneNumber string     `gorm:"column:phone_number"`
	Address     string     `gorm:"column:address"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

// TableName returns the table name for the customers model
func (*Customers) TableName() string {
	return "customers"
}
