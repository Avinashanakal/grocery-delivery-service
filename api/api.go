// Package api provides interface between the server data exchange & business need
package api

import (
	"context"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/groceries"
	"github.com/Avinashanakal/grocery-delivery-service/internal/pkg/models"
	"github.com/Avinashanakal/grocery-delivery-service/platform/config"
)

type API struct {
	config config.Config
	db     *gorm.DB
	log    *logrus.Logger

	groceries groceries.Service
}

func New(config config.Config, db *gorm.DB, log *logrus.Logger) API {
	return API{
		config:    config,
		db:        db,
		log:       log,
		groceries: groceries.New(db, log),
	}
}

func (api *API) Groceries(ctx context.Context, grocery groceries.GroceryRequest) error {

	return api.groceries.Groceries(ctx, grocery)
}

func (api *API) Search(ctx context.Context,
	searchQuery string,
) ([]groceries.GrocerySearchResult, error) {

	return api.groceries.Search(ctx, searchQuery)
}

func (api *API) Buy(ctx context.Context, orderRequest groceries.OrderRequest) (models.Order, error) {
	return api.groceries.Buy(ctx, orderRequest)
}

func (api *API) Orders(ctx context.Context, customerID int) ([]groceries.OrderResponse, error) {
	return api.groceries.Orders(ctx, customerID)
}
