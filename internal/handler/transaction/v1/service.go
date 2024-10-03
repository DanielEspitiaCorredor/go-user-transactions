package transactionv1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/DanielEspitiaCorredor/go-user-transactions/tools"
)

func init() {

}

func GenerateReport(ctx *gin.Context) {

	var request ExtractRequest
	requestID := uuid.NewString()
	msg := fmt.Sprintf("[GenerateReport][RequestID:%s]", requestID)

	// Bind request data
	if errMsg, err := tools.BindRequestData(ctx, &request); err != nil {
		fmt.Println(msg, "BindRequestData Error", err)
		tools.SendError(ctx, nil, err, http.StatusBadRequest, errMsg)
		return
	}

	if isValidEmail := tools.ValidateEmail(request.GetReceiverEmail()); !isValidEmail {
		fmt.Println(msg, "Invalid receiver email", request.GetReceiverEmail())
		tools.SendError(ctx, nil, errors.New("invalid params"), http.StatusBadRequest, "no valid receiver_email")
		return

	}

	msg = fmt.Sprintf("%s[Account:%s][Year:%d]", msg, request.GetAccount(), request.GetYear())

	fmt.Println(msg)

	txDf, err := NewTransactionDF(request.GetAccount(), fmt.Sprintf("./assets/account/account_%s.csv", request.GetAccount()), request.GetYear())

	if err != nil {

		fmt.Println(msg, "NewTransactionDF err", err)
		tools.SendError(ctx, nil, err, http.StatusFailedDependency, "error processing file")
		return
	}

	// Pre Process file data
	txDf.PreProcessData()
	// Get account balance
	balance := txDf.NewAccountBalance()
	balance.SendReport(request.GetReceiverEmail())

	tools.SendResponse(ctx, http.StatusNoContent, nil, nil, tools.GinResponseTypes_JSON)

}
