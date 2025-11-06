/*
docker run --name myapp-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=bank_account_api_db \
  -p 5432:5432 \
  -d postgres:16
*/

/*
psql -U admin -d bank_account_api_db
*/

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
