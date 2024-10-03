package tools

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type GinResponseType int32

const (
	GinResponseTypes_UNKNOWN GinResponseType = iota
	GinResponseTypes_JSON
	GinResponseTypes_NOCONTENT
)

func SendResponse(ctx *gin.Context, statusCode int, data any, headers map[string]string, ginRespType GinResponseType) (err error) {

	for k, v := range headers {

		if strings.ToLower(k) == "content-type" {
			continue
		}

		ctx.Header(k, v)

	}

	switch ginRespType {
	case GinResponseTypes_JSON:

		ctx.JSON(statusCode, data)

	case GinResponseTypes_NOCONTENT:

		ctx.Status(statusCode)
	default:

		err = errors.New("invalid gin response type")
	}

	return
}

/*
Output an error message
*/
func SendError(ctx *gin.Context, response interface{}, err interface{}, statusCode int, messages ...string) {
	message := ""
	for _, m := range messages {
		if message == "" {
			message = m
		} else {
			message = fmt.Sprintf("%s - %s", message, m)
		}
	}

	ctx.JSON(statusCode, gin.H{
		"response": response,
		"message":  message,
		"error":    fmt.Sprint(err),
	})

}
