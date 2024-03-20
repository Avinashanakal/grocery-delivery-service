package groceries

import (
	"context"
	"fmt"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/models"
)

/*
TODO
- we can move db calls to separate package for better reaadability
*/
func (s *service) Orders(ctx context.Context, customerID int) ([]OrderResponse, error) {

	var orders []models.Order
	if err := s.db.Where("customer_id = ? AND status = ?", customerID, "PLACED").Find(&orders).Error; err != nil {
		s.log.WithError(err).Error("Failed to fetch orders")
		return []OrderResponse{}, err
	}

	// Check if no orders were found
	if len(orders) == 0 {
		notFoundErr := fmt.Errorf("No orders found for customer with ID: %d", customerID)
		s.log.Error(notFoundErr.Error())
		return []OrderResponse{}, notFoundErr
	}

	var orderResponses []OrderResponse
	for _, order := range orders {
		orderResponse := OrderResponse{
			ID:         order.ID,
			CustomerID: order.CustomerID,
			Status:     string(order.Status),
			Price:      order.TotalPrice,
		}

		// Find all order items for the order.
		var orderItems []models.OrderItems
		if err := s.db.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			s.log.WithError(err).Error("Failed to find Order Items")
			return []OrderResponse{}, err
		}

		// Convert the order items to a slice of order item responses.
		var orderItemResponses []OrderItemResponse
		for _, orderItem := range orderItems {
			orderItemResponse := OrderItemResponse{
				ID:        orderItem.ID,
				OrderID:   orderItem.OrderID,
				GroceryID: orderItem.GroceryID,
				Quantity:  orderItem.Quantity,
			}

			orderItemResponses = append(orderItemResponses, orderItemResponse)
		}

		// Assign orderItemResponses only if there are order items
		if len(orderItemResponses) > 0 {
			orderResponse.OrderItems = orderItemResponses
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses, nil
}
