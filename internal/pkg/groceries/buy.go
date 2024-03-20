package groceries

import (
	"context"
	"time"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/models"
)

// TODO handle ignoring timestamp with TZ in the below response
func (s *service) Buy(ctx context.Context, orderRequest OrderRequest) (models.Order, error) {

	order := models.Order{
		CustomerID:   orderRequest.CustomerID,
		Status:       models.Placed,
		OrderDate:    time.Now(),
		DeliveryDate: time.Now().Add(time.Hour * 1), // As of now added +1 hour for delivery date
	}

	//add  order for that customer
	if err := s.db.Create(&order).Error; err != nil {
		s.log.WithError(err).Error("Failed to add new order in the Order table")
		return models.Order{}, err
	}

	var totalPrice float64

	// adding order items for each grocery item in the request
	for _, orderItemRequest := range orderRequest.OrderItems {
		orderItem := models.OrderItems{
			OrderID:   order.ID,
			GroceryID: orderItemRequest.GroceryID,
			Quantity:  orderItemRequest.Quantity,
		}

		if err := s.db.Create(&orderItem).Error; err != nil {
			s.log.WithError(err).Error("Failed to add new order in the Order table")
			return models.Order{}, err
		}

		//Fetch price for the grocery id from grocery table
		var groceryPrice float64
		if err := s.db.Model(&models.Grocery{}).
			Select("price").
			Where("id=?", orderItemRequest.GroceryID).
			First(&groceryPrice).Error; err != nil {
			s.log.WithError(err).Error("Failed to fetch grocery price")
		}

		totalPrice += groceryPrice

	}
	order.TotalPrice = totalPrice

	s.log.Info("Successfully Placed the grocery items")

	return order, nil
}
