package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zweimach/xendit-trial/domain"
	"github.com/zweimach/xendit-trial/utils"
)

type AuditTransactionsRequest struct {
	A []domain.Transaction `json:"a" validate:"required"`
	B []domain.Transaction `json:"b" validate:"required"`
}

type AuditTransactionsResponse struct {
	PaymentRefID string `json:"payment_ref_id"`
	AuditResult  string `json:"audit_result"`
}

func AuditTransactions(c echo.Context) error {
	errRes := utils.NewError()

	req := new(AuditTransactionsRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	result := make([]AuditTransactionsResponse, 10)

	return c.JSON(http.StatusOK, result)
}
