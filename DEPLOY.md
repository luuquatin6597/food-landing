# Render Deployment Guide

## Option 1: Using render.yaml (Recommended)

The project includes a `render.yaml` file that automatically configures the deployment:

1. Connect your GitHub repository to Render
2. Select "Use YAML file"
3. Render will automatically read `render.yaml` and configure:
   - Root directory: `backend`
   - Build command: `go build -o server main.go`
   - Start command: `./server`

## Option 2: Manual Configuration

### Backend Service Setup:

1. **Service Type**: Web Service
2. **Repository**: Connect your GitHub repo
3. **Root Directory**: `backend`
4. **Runtime**: Go
5. **Build Command**: `go build -o server main.go`
6. **Start Command**: `./server`

### Environment Variables:

- `DATABASE_URL`: Connect to your PostgreSQL database
- `PORT`: Automatically set by Render (usually 10000)

### Database Setup:

1. Create PostgreSQL database service
2. Note the connection string
3. Use it in `DATABASE_URL` environment variable

## Testing Deployment

After deployment, test your API:

- `GET https://your-app.onrender.com/api/foods`

## Important Notes

- The app automatically creates tables and seeds data on first run
- Make sure your PostgreSQL database is running before the web service
- The app includes CORS headers for frontend integration
