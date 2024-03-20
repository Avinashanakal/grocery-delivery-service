package http

import "github.com/Avinashanakal/grocery-delivery-service/api"

type Controller struct {
	api api.API
}

func NewController(api api.API) Controller {
	return Controller{
		api: api,
	}
}
