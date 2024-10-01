package tools

import "github.com/gin-gonic/gin"

// Bind request data on interface object
func BindRequestData(ctx *gin.Context, dstObj interface{}) (string, error) {
	if err := ctx.Bind(dstObj); err != nil {
		return "Error binding querystring params", err
	}

	if err := ctx.BindHeader(dstObj); err != nil {
		return "Error binding headers", err
	}
	// Add more binds if you needed
	return "", nil
}
