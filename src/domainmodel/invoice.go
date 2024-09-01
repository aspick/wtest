package domainmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	// 手数料率 (パーセント)
	ChargeRatePercentage = 4
	// 消費税率 (パーセント)
	TaxRatePercentage = 10
)

// Invoice は請求書を表すドメインモデルです。
type Invoice struct {
	CompanyID      int
	CustomerID     int
	IssueDate      time.Time
	PaymentAmount  decimal.Decimal
	Charge         decimal.Decimal
	ChargeRate     decimal.Decimal
	ConsumptionTax decimal.Decimal
	BillingAmount  decimal.Decimal
	PaymentDueDate time.Time
	Status         InvoiceStatus
}

type InvoiceStatus string

const (
	InvoiceStatusPending    InvoiceStatus = "pending"
	InvoiceStatusProcessing InvoiceStatus = "processing"
	InvoiceStatusCompleted  InvoiceStatus = "completed"
	InvoiceStatusFailed     InvoiceStatus = "failed"
)

// NewInvoice は,請求書を作成します。
// ここで、手数料や消費税、請求金額を計算します。
func NewInvoice(
	companyID int,
	customerID int,
	issueDate time.Time,
	paymentAmount decimal.Decimal,
	paymentDueDate time.Time,
) Invoice {
	// 手数料率は 4%
	chargeRate := decimal.NewFromFloat(ChargeRatePercentage).Div(decimal.NewFromFloat(100))
	// 消費税率は 10%
	taxRate := decimal.NewFromFloat(TaxRatePercentage).Div(decimal.NewFromFloat(100))

	// 手数料を計算
	charge := paymentAmount.Mul(chargeRate)

	// 消費税を計算
	tax := charge.Mul(taxRate)

	// 請求金額を計算
	billingAmount := paymentAmount.Add(charge).Add(tax)

	return Invoice{
		CompanyID:      companyID,
		CustomerID:     customerID,
		IssueDate:      issueDate,
		PaymentAmount:  paymentAmount,
		Charge:         charge,
		ChargeRate:     chargeRate,
		ConsumptionTax: tax,
		BillingAmount:  billingAmount,
		PaymentDueDate: paymentDueDate,
		Status:         InvoiceStatusPending,
	}
}
