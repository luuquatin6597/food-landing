package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"food-landing-backend/database"
	"food-landing-backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Render sẽ cung cấp biến môi trường này tự động
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Chạy hàm migrate và seed dữ liệu
	if err := database.MigrateAndSeed(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/foods", getFoods)

	// Render yêu cầu chạy trên port 10000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Chạy port 8080 ở local
	}
	router.Run(":" + port)
}

func getFoods(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, description, ingredients, price, image_url, color, region FROM foods ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	foods := []models.Food{}
	for rows.Next() {
		var f models.Food
		if err := rows.Scan(&f.ID, &f.Name, &f.Description, &f.Ingredients, &f.Price, &f.ImageURL, &f.Color, &f.Region); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		foods = append(foods, f)
	}
	c.JSON(http.StatusOK, foods)
}
