INSERT INTO companies (name, representative_name, telephone, zip_code, address) VALUES
  ('株式会社A', '山田太郎', '090-1234-5678', '123-4567', '東京都千代田区1-1-1')
;

INSERT INTO users (company_id, name, email, hashed_password) VALUES
  (1, '山田太郎', 'test@example.com', 'test-hashed-password')
;

INSERT INTO customers (company_id, name, representative_name, telephone, zip_code, address) VALUES
  (1, '株式会社B', '田中次郎', '090-1234-5678', '123-4567', '東京都千代田区1-1-1')
;

INSERT INTO customer_band_accounts (company_id, customer_id, bank_name, branch_name, account_number, account_name) VALUES
  (1, 1, '銀行A', '支店A', '1234567', '株式会社B')
;
