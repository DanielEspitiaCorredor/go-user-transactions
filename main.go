package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/DanielEspitiaCorredor/go-user-transactions/internal/middleware"
	"github.com/DanielEspitiaCorredor/go-user-transactions/internal/routes"
)

func setupRouter() *gin.Engine {

	// General router configuration
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ValidateApiKey("SERVICE_API_KEY"))

	// Router map
	svcRouters := &gin.RouterGroup{}
	healthzResponse := gin.H{
		"message": "Hello from go-user-transactions",
	}

	serviceVersion := routes.SvcVersion(os.Getenv("SERVICE_VERSION"))

	switch serviceVersion {

	case routes.ServiceVersion_V1:

		svcRouters = r.Group("api/v1")

		healthzResponse = gin.H{
			"status":  "running",
			"message": "Hello from go-user-transactions",
			"version": serviceVersion,
		}

	default:
		panic(fmt.Sprintf("version '%s' not implemented", serviceVersion))
	}

	// Health service
	svcRouters.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthzResponse)
	})

	// routes
	routes.MapTransactionRoutes(svcRouters, serviceVersion)

	return r
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	port := os.Getenv("SERVICE_PORT")
	if port == "" {

		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
}
