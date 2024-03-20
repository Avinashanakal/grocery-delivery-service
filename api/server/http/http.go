package http

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/Avinashanakal/grocery-delivery-service/api"
	"github.com/Avinashanakal/grocery-delivery-service/api/server/http/middleware"
)

func New(api api.API) *gin.Engine {
	// Check if the environment is production
	isProdEnv := os.Getenv("TIER") == "production"

	// Set Gin mode based on the environment
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// Use Gin logger middleware in debug mode
	if !isProdEnv {
		gin.SetMode(gin.DebugMode)
		engine.Use(gin.Logger())
	}

	// Add the secret-based authentication middleware
	engine.Use(middleware.SecretAuthenticationMiddleware)

	// Handle requests with no matching route with a 404 status
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	// Define a simple health check endpoint
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	ctrl := NewController(api)
	path := "/v1/grocery"
	v1 := engine.Group(path)
	{
		v1.POST("/", ctrl.Groceries)
		v1.GET("/search", ctrl.Search)
		v1.POST("/buy", ctrl.Buy)
		v1.GET("/orders/:customerID", ctrl.Orders)
	}
	return engine
}
