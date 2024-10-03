package routes

import (
	"github.com/gin-gonic/gin"

	transactionv1 "github.com/DanielEspitiaCorredor/go-user-transactions/internal/handler/transaction/v1"
)

func MapTransactionRoutes(router *gin.RouterGroup, version SvcVersion) {

	txGroup := router.Group("/transactions")

	switch version {

	case ServiceVersion_V1:

		// endpoint used to generate reports getting data from csv file
		txGroup.POST("/generate_report", transactionv1.GenerateReport)
	}

}
