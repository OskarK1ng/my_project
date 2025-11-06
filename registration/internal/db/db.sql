/*
docker run --name myapp-postgres \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=my_app_db \
  -p 5432:5432 \
  -d postgres:16
*/

/*
psql -U postgres -d mydb --pset pager=off
*/

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id UUID UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    user_last_name VARCHAR(50) NOT NULL,
    user_email VARCHAR(50) UNIQUE NOT NULL,
    user_phone VARCHAR(20) UNIQUE NOT NULL,
    user_password_hash TEXT NOT NULL,
    user_created_at TIMESTAMP DEFAULT NOW()
);
