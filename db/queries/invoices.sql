-- name: CreateInvoice :one
INSERT INTO invoices (
  id,
  company_id,
  customer_id,
  issue_date,
  payment_amount,
  charge,
  charge_rate,
  consumption_tax,
  billing_amount,
  payment_due_date,
  status
)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;
