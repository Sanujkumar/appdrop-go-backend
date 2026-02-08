# AppDrop Mini App Config API

REST API built in Go for managing mobile app pages and widgets.
This project is a simplified backend for AppDrop’s no-code mobile app builder.

---

## 🚀 Tech Stack

- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL (Neon DB)
- UUID
- JSONB (for widget config storage)

---

## 📦 Features

✅ Create, update, delete Pages  
✅ Add, update, delete Widgets  
✅ Reorder widgets on a page  
✅ PostgreSQL JSONB support  
✅ Validation rules implemented  
✅ Consistent error responses  

---

## 📁 Project Structure

appdrop-backend/
│
├── main.go
├── go.mod
├── .env
├── middleware/
│   └── logger.go
│
├── config/
│ └── database.go
│
├── models/
│ ├── page.go
│ └── widget.go
│
├── handlers/
│ ├── page_handler.go
│ └── widget_handler.go
│
├── routes/
│ └── routes.go
│
└── utils/
└── response.go


---

## ⚙️ Setup Instructions

### 1️⃣ Install Go
Download from:
https://go.dev/dl/

Verify installation:
go version


---

### 2️⃣ Clone Repository
git clone <your-repository-url>
cd appdrop-backend


---

### 3️⃣ Install Dependencies
go mod tidy


---

### 4️⃣ Setup Environment Variables

Create `.env` file in root folder:

DATABASE_URL=your_neon_postgres_connection_string
PORT=8008


Example Neon connection string:
postgres://user:password@host/dbname?sslmode=require


---

### 5️⃣ Run Server
go run main.go


Server runs at:
http://localhost:8008


Health check:
GET http://localhost:8008/health


---

## 🗄️ Database Setup

Database tables are automatically created using GORM AutoMigrate.

### Page Table
- id (UUID)
- name (string, required)
- route (unique string, required)
- is_home (boolean)
- created_at
- updated_at

### Widget Table
- id (UUID)
- page_id (UUID, foreign key)
- type (string)
- position (integer)
- config (JSONB)
- created_at
- updated_at

---

## 📡 API Endpoints

### Pages

#### GET all pages
GET /pages


#### GET page with widgets
GET /pages/:id


#### Create page
POST /pages
Content-Type: application/json

{
"name": "Home",
"route": "/home",
"is_home": true
}


#### Update page
PUT /pages/:id


#### Delete page
DELETE /pages/:id


---

### Widgets

#### Add widget to page
POST /pages/:id/widgets
Content-Type: application/json

{
"type": "banner",
"position": 1,
"config": {
"image_url": "https://example.com/banner.jpg"
}
}


Allowed widget types:
- banner
- product_grid
- text
- image
- spacer

#### Update widget
PUT /widgets/:id


#### Delete widget
DELETE /widgets/:id


#### Reorder widgets
POST /pages/:id/widgets/reorder
Content-Type: application/json

[
"widget_id_1",
"widget_id_2",
"widget_id_3"
]


---

## ❌ Error Response Format

{
"error": {
"code": "VALIDATION_ERROR",
"message": "Page route already exists"
}
}


---

## 🧪 Example API Requests (curl)

### Create Page
curl -X POST http://localhost:8008/pages
-H "Content-Type: application/json"
-d '{"name":"Home","route":"/home","is_home":true}'


### Get All Pages
curl http://localhost:8008/pages


### Add Widget
curl -X POST http://localhost:8008/pages/<PAGE_ID>/widgets
-H "Content-Type: application/json"
-d '{"type":"banner","position":1,"config":{"image_url":"https://example.com/banner.jpg"}}'


---

## ✅ Validation Rules Implemented

- Page route must be unique
- Page name is required
- Only one home page allowed
- Cannot delete home page
- Widget type must be valid
- Widget config must be valid JSON

---

Bonus Features:
- Pagination for GET /pages
- Widget filtering by type
- Request logging middleware


## 👨‍💻 Author

Sanuj_Kumar