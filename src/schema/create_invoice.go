package schema

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateInvoiceRequest struct {
	// 顧客ID
	CustomerID int `json:"customer_id"`
	// 発行日
	IssueDate time.Time `json:"issue_date"`
	// 請求金額
	PaymentAmount decimal.Decimal `json:"payment_amount"`
	// 支払期日
	PaymentDueDate time.Time `json:"due_date"`
}

type CreateInvoiceResponse struct {
	// 請求書ID
	InvoiceID int `json:"invoice_id"`
	// TODO: 契約書の情報を追加する
}
