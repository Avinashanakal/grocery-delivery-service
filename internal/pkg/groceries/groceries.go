package groceries

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/models"
)

type service struct {
	db  *gorm.DB
	log *logrus.Logger
}

func New(db *gorm.DB, log *logrus.Logger) Service {
	return &service{
		db:  db,
		log: log,
	}
}

type Service interface {
	Groceries(ctx context.Context, grocery GroceryRequest) error

	Search(ctx context.Context,
		searchQuery string,
	) ([]GrocerySearchResult, error)

	Buy(ctx context.Context, orderRequest OrderRequest) (models.Order, error)

	Orders(ctx context.Context, customerID int) ([]OrderResponse, error)
}

func (s *service) Groceries(ctx context.Context, grocery GroceryRequest) error {

	groceryData := models.Grocery{
		Name:        grocery.Name,
		Description: grocery.Description,
		Price:       grocery.Price,
		Quantity:    grocery.Quantity,
	}

	err := s.db.Create(&groceryData).Error
	if err != nil {
		s.log.WithError(err).Error("Failed to add groceries to inventory")
		return err
	}

	s.log.Info("Successfully added groceries to inventory")

	return nil
}
