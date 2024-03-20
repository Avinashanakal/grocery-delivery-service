package groceries

type GroceryRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
}

type GrocerySearchResult struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type OrderRequest struct {
	CustomerID int64              `json:"customer_id" binding:"required"`
	OrderItems []OrderItemRequest `json:"order_items" binding:"required"`
}

type OrderItemRequest struct {
	GroceryID int64 `json:"grocery_id" binding:"required"`
	Quantity  int   `json:"quantity" binding:"required"`
}

type OrderResponse struct {
	ID         int                 `json:"id"`
	CustomerID int64               `json:"customer_id"`
	OrderItems []OrderItemResponse `json:"order_items,omitempty"`
	Status     string              `json:"status"`
	Price      float64             `json:"price"`
}

type OrderItemResponse struct {
	ID        int   `json:"id"`
	OrderID   int   `json:"order_id"`
	GroceryID int64 `json:"grocery_id"`
	Quantity  int   `json:"quantity"`
}
