package usecase

import (
	"context"

	"github.com/aspick/wtest/src/db"
	"github.com/aspick/wtest/src/domainmodel"
	"github.com/aspick/wtest/src/schema"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

// CreateInvoice は請求書を作成するユースケースです。
type CreateInvoice interface {
	Execute(ctx context.Context, companyID int, request schema.CreateInvoiceRequest) (schema.CreateInvoiceResponse, error)
}

type createInvoiceImpl struct {
	pool *pgxpool.Pool
}

func NewCreateInvoice(pool *pgxpool.Pool) CreateInvoice {
	return &createInvoiceImpl{pool: pool}
}

// Execute は請求書を作成します。
func (c *createInvoiceImpl) Execute(ctx context.Context, companyID int, request schema.CreateInvoiceRequest) (schema.CreateInvoiceResponse, error) {
	invoice := domainmodel.NewInvoice(
		companyID,
		request.CustomerID,
		request.IssueDate,
		request.PaymentAmount,
		request.PaymentDueDate,
	)
	// TODO: ドメインモデルの Invoice の validation

	// TODO:: 以下の処理を repository に移動する
	invoiceDBModel, err := c.convertDomainmodelToDBModel(invoice)
	if err != nil {
		return schema.CreateInvoiceResponse{}, err
	}

	query := db.New(c.pool)
	created, err := query.CreateInvoice(ctx, invoiceDBModel)
	if err != nil {
		return schema.CreateInvoiceResponse{}, err
	}

	return schema.CreateInvoiceResponse{
		InvoiceID: int(created.ID),
	}, nil
}

func (c *createInvoiceImpl) convertDomainmodelToDBModel(invoice domainmodel.Invoice) (db.CreateInvoiceParams, error) {
	paymentAmount, err := convertDecimalToPgtypeNumeric(invoice.PaymentAmount)
	if err != nil {
		return db.CreateInvoiceParams{}, err
	}

	charge, err := convertDecimalToPgtypeNumeric(invoice.Charge)
	if err != nil {
		return db.CreateInvoiceParams{}, err
	}

	chargeRate, err := convertDecimalToPgtypeNumeric(invoice.ChargeRate)
	if err != nil {
		return db.CreateInvoiceParams{}, err
	}

	consumptionTax, err := convertDecimalToPgtypeNumeric(invoice.ConsumptionTax)
	if err != nil {
		return db.CreateInvoiceParams{}, err
	}

	billingAmount, err := convertDecimalToPgtypeNumeric(invoice.BillingAmount)
	if err != nil {
		return db.CreateInvoiceParams{}, err
	}

	return db.CreateInvoiceParams{
		CompanyID:      int32(invoice.CompanyID),
		CustomerID:     int32(invoice.CustomerID),
		IssueDate:      pgtype.Date{Time: invoice.IssueDate, Valid: true},
		PaymentAmount:  paymentAmount,
		Charge:         charge,
		ChargeRate:     chargeRate,
		ConsumptionTax: consumptionTax,
		BillingAmount:  billingAmount,
		PaymentDueDate: pgtype.Date{Time: invoice.PaymentDueDate, Valid: true},
		Status:         db.InvoiceStatus(invoice.Status),
	}, nil
}

func convertDecimalToPgtypeNumeric(decimal decimal.Decimal) (pgtype.Numeric, error) {
	numeric := pgtype.Numeric{}
	err := numeric.Scan(decimal.String())
	if err != nil {
		return pgtype.Numeric{}, err
	}
	return numeric, nil
}
