/*
docker run --name myapp-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=my_app_db \
  -p 5432:5432 \
  -d postgres:16
*/

/*
psql -U admin -d my_app_db
psql -U admin -d my_app_db --pset pager=off 
*/

-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL, 
    user_name VARCHAR(50) NOT NULL,
    user_last_name VARCHAR(50) NOT NULL,
    user_email VARCHAR(50) UNIQUE NOT NULL,
    user_phone VARCHAR(20) UNIQUE NOT NULL,
    user_password_hash TEXT NOT NULL,
    user_created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Таблица транзакций / счетов
CREATE TABLE IF NOT EXISTS balance (
    user_id UUID NOT NULL UNIQUE REFERENCES users(user_id) ON DELETE CASCADE,
    balance NUMERIC(15,2) NOT NULL DEFAULT 0
);
