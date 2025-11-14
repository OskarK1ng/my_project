# 🧩 My Project — User Registration, Authentication & Balance Service

## 📘 Описание

Этот проект реализует базовую микросервисную архитектуру на Go с использованием **PostgreSQL** и **Docker Compose**.  
Он включает три основных компонента:

1. **Registration Service** — регистрация пользователей.  
   После создания нового пользователя автоматически создаётся счёт (balance = 0).
2. **Auth Service** — авторизация и выдача JWT-токенов.
3. **Transactions / Balance Service** — работа с балансом пользователя (пополнение и снятие).

---

## ⚙️ Технологии

- **Golang 1.22+**
- **PostgreSQL 15**
- **Docker / Docker Compose**
- **pgx v5**
- **bcrypt** для хэширования паролей
- **JWT (JSON Web Token)** для аутентификации
- Архитектура: **Clean Architecture / Layered (Handlers → Services → Repository → DB)**

---

## 📂 Архитектура проекта

```
my_project/
│
├─ auth/ # Модуль аутентификации
│ ├─ cmd/
│ │ └─ app/
│ │   └─ main.go
│ │
│ └─ internal/
│   ├─ config/
│   │ └─ config.go
│   │
│   ├─ db/
│   │ └─ db.go
│   │
│   ├─ handlers/
│   │ └─ auth.go
│   │
│   ├─ models/
│   │ └─ user.go
│   │
│   ├─ security/
│   │ └─ jwt.go
│   │
│   ├─ services/
│   │  └─ auth.go
│   │
│   ├─ .env
│   ├─ Dockerfile
│   ├─ go.mod
│   └─ go.sum
├─ registration/ # Модуль регистрации
│ ├─ cmd/
│ │ └─ app/
│ │   └─ main.go
│ │
│ └─ internal/
│   ├─ config/
│   │ └─ config.go
│   │
│   ├─ db/
│   │ └─ db.go
│   │
│   ├─ handlers/
│   │ └─ register.go
│   │
│   ├─ models/
│   │ └─ user.go
│   │
│   ├─ repository/
│   │ └─ balance.go
│   │
│   ├─ services/
│   │  └─ register.go
│   │
│   ├─ .env
│   ├─ Dockerfile
│   ├─ go.mod
│   └─ go.sum
├─ transactions/ # Модуль пополнения и снятия
│ ├─ cmd/
│ │ └─ app/
│ │   └─ main.go
│ │
│ └─ internal/
│   ├─ config/
│   │ └─ config.go
│   │
│   ├─ db/
│   │ └─ db.go
│   │
│   ├─ handlers/
│   │ └─ handler.go
│   │
│   ├─ middleware/
│   │ └─ auth.go
│   │
│   ├─ models/
│   │ └─ account.go
│   │
│   ├─ repository/
│   │ └─ balance.go
│   │
│   ├─ security/
│   │ └─ jwt.go
│   │
│   ├─ services/
│   │  └─ balance.go
│   │
│   ├─ .env
│   ├─ Dockerfile
│   ├─ go.mod
│   └─ go.sum
├── .gitignore
├── db.sql
├── docker-compose.yml
└── README.md
```

---

## 📦 Установка и запуск

### Клонируйте репозиторий

```bash
git clone https://github.com/OskarK1ng/my_project.git
cd my_project
```

### Настройте .env

#### /auth
```sh
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=my_app_db
POSTGRES_HOST=db
POSTGRES_PORT=5432
SERVER_PORT=:8082

JWT_SECRET=super_secret_key_123
JWT_TTL_MINUTES=15
```
#### /registration
```sh
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=my_app_db
POSTGRES_HOST=db
POSTGRES_PORT=5432
SERVER_PORT=:8081

JWT_SECRET=super_secret_key_123
JWT_TTL_MINUTES=15
```
#### /transactions
```sh
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=my_app_db
POSTGRES_HOST=db
POSTGRES_PORT=5432
SERVER_PORT=:8080

JWT_SECRET=super_secret_key_123
JWT_TTL_MINUTES=15
```

### Запустите Docker Compose

```bash
docker-compose down -v  # удалить старые контейнеры и данные
docker-compose up --build -d
```

### Postman

Вы можете протестировать API с помощью Postman. Все эндпоинты собраны в коллекцию:

[Перейти в коллекцию](https://identity.getpostman.com/login?addAccount=1&email=askarov.30040%40gmail.com&user_id=47498884&action_performed=accountSelect&authFlowId=e72da1ee-2fa2-4d5e-a2b6-40edca4e4c6a)

