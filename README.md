# Food Landing - Vietnamese Food API

A full-stack application showcasing Vietnamese cuisine with Go backend and React frontend.

## 🚀 Features

- **Backend**: Go with Gin framework
- **Database**: PostgreSQL with auto-migration
- **Frontend**: React with Axios
- **Deployment**: Ready for Render deployment

## 🛠️ Local Development

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Node.js 18+

### Backend Setup

1. Start PostgreSQL:

```bash
docker-compose up -d
```

2. Start backend server:

```bash
cd backend
export DATABASE_URL="postgres://fooduser:foodpassword@localhost:5433/fooddb?sslmode=disable"
export PORT=8080
go run main.go
```

### Frontend Setup

```bash
cd frontend/frontend
npm install
npm start
```

## 🌐 API Endpoints

- `GET /api/foods` - Get all Vietnamese food items

## 🚀 Deployment on Render

### Backend Service

1. Connect your GitHub repository
2. Set root directory: `backend`
3. Set build command: `go build -o server main.go`
4. Set start command: `./server`
5. Add environment variables:
   - `DATABASE_URL` (from PostgreSQL database)
   - `PORT` (automatically set by Render)

### Database

1. Create PostgreSQL database on Render
2. Use the connection string in your backend service

## 📁 Project Structure

```
food-landing/
├── backend/             # Go backend API
│   ├── main.go         # Main application entry
│   ├── go.mod          # Go module file
│   ├── database/       # Database migrations & seeds
│   └── models/         # Data models
├── frontend/           # React frontend
├── docker-compose.yml  # Local PostgreSQL setup
└── render.yaml         # Render deployment config
```

## 🍜 Sample Data

The application includes 9 Vietnamese dishes:

- Mì Quảng, Cao Lầu, Bánh Mì, Gỏi Cuốn
- Bánh Xèo, Phở, Bún Bò Huế, Bánh Canh, Cơm Gà Tam Kỳ

## 🔧 Environment Variables

- `DATABASE_URL`: PostgreSQL connection string
- `PORT`: Server port (default: 8080, Render uses 10000)
