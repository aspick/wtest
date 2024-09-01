CREATE TABLE test_table (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- 企業
CREATE TABLE companies ( 
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL, -- 法人名
  representative_name VARCHAR(255) NOT NULL,-- 代表者名
  telephone VARCHAR(20) NOT NULL, -- 電話番号
  zip_code VARCHAR(10) NOT NULL, -- 郵便番号
  address VARCHAR(255) NOT NULL -- 住所
);

-- ユーザー
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  company_id INTEGER NOT NULL, -- 企業ID
  name VARCHAR(255) NOT NULL, -- ユーザー名
  email VARCHAR(255) NOT NULL, -- メールアドレス
  hashed_password VARCHAR(255) NOT NULL, -- パスワード

  FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- 取引先
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  company_id INTEGER NOT NULL, -- 企業ID
  name VARCHAR(255) NOT NULL, -- 取引先名
  representative_name VARCHAR(255) NOT NULL, -- 代表者名
  telephone VARCHAR(20) NOT NULL, -- 電話番号
  zip_code VARCHAR(10) NOT NULL, -- 郵便番号
  address VARCHAR(255) NOT NULL, -- 住所

  FOREIGN KEY (company_id) REFERENCES companies(id) 
);

-- 取引先銀行口座
CREATE TABLE customer_band_accounts (
  id SERIAL PRIMARY KEY,
  company_id INTEGER NOT NULL, -- 企業ID
  customer_id INTEGER NOT NULL, -- 取引先ID
  bank_name VARCHAR(255) NOT NULL, -- 銀行名
  branch_name VARCHAR(255) NOT NULL, -- 支店名
  account_number VARCHAR(20) NOT NULL, -- 口座番号
  account_name VARCHAR(255) NOT NULL, -- 口座名

  FOREIGN KEY (company_id) REFERENCES companies(id),
  FOREIGN KEY (customer_id) REFERENCES customers(id)
);


-- 請求書のステータス
CREATE TYPE invoice_status AS ENUM (
  'pending', -- 未処理
  'processing', -- 処理中
  'completed', -- 支払済み
  'failed' -- エラー
);

-- 請求書データ
CREATE TABLE invoices (
  id SERIAL PRIMARY KEY,
  company_id INTEGER NOT NULL, -- 企業ID
  customer_id INTEGER NOT NULL, -- 取引先ID
  issue_date DATE NOT NULL, -- 発行日
  payment_amount decimal NOT NULL, -- 支払金額
  charge decimal NOT NULL, -- 手数料
  charge_rate decimal(8,7) NOT NULL, -- 手数料率
  consumption_tax decimal NOT NULL, -- 消費税
  billing_amount decimal NOT NULL, -- 請求金額
  payment_due_date DATE NOT NULL, -- 支払期限
  status invoice_status NOT NULL, -- ステータス

  FOREIGN KEY (company_id) REFERENCES companies(id),
  FOREIGN KEY (customer_id) REFERENCES customers(id)
); 
