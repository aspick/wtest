# web api

## requirements

- go 1.23
- docker
- docker-compose
- sqlc
- direnv

## API schema

### Create Invoice

`POST /api/invoices`

request body
```json
{
  "customer_id": 1, // 取引先ID
  "issue_date": "2024-01-01", // 発行日
  "payment_amount": "1000000", // 支払い金額
  "payment_due_date": "2024-01-31", // 支払い期限
}
```

response body
```json
{
  "id": 1, // 請求書ID
  // その他の属性は今後追加
}
```

### Get Invoices

`GET /api/invoices`

request params

TBD

response body

TBD
