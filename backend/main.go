package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Food struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Price       string `json:"price"`
	ImageURL    string `json:"image_url"`
	Color       string `json:"color"`
	Region      string `json:"region"`
}

var db *sql.DB

// Hàm mới để tự động tạo bảng và chèn dữ liệu
func migrateAndSeed(db *sql.DB) {
	// 1. Tạo bảng nếu chưa tồn tại
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS foods (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        ingredients TEXT,
        price VARCHAR(100),
        image_url VARCHAR(255),
        color VARCHAR(20),
        region VARCHAR(100)
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Không thể tạo bảng: %v", err)
	}

	// 2. Kiểm tra xem bảng có dữ liệu chưa
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM foods").Scan(&count)
	if err != nil {
		log.Fatalf("Không thể đếm dữ liệu: %v", err)
	}

	// 3. Nếu bảng trống, chèn dữ liệu
	if count == 0 {
		log.Println("Bảng foods trống, bắt đầu chèn dữ liệu...")
		insertSQL := `
        INSERT INTO foods (name, description, ingredients, price, image_url, color, region) VALUES
        ('Mì Quảng', 'Mì Quảng là món ăn đặc sản nổi tiếng nhất của Quảng Nam với nước dùng đậm đà từ tôm cua, thịt heo, trứng cút và bánh tráng nướng thơm ngon.', 'Mì vàng, Tôm tươi, Thịt heo, Trứng cút, Bánh tráng nướng, Rau thơm', '35.000 - 50.000 VNĐ', 'https://danangfantasticity.com/wp-content/uploads/2024/04/cach-thuong-thuuc-mot-to-mi-quang-dung-dieu-nguoi-da-nang.jpg', '#ff6b35', 'Quảng Nam'),
        ('Cao Lầu', 'Cao Lầu là món ăn đặc sản của Hội An, với sợi mì dày, thịt xá xíu, rau sống và nước dùng đậm đà.', 'Mì cao lầu, Thịt xá xíu, Rau sống, Nước dùng', '30.000 - 45.000 VNĐ', 'https://i-giadinh.vnecdn.net/2023/03/13/Buoc-7-Thanh-pham-1-7-9577-1678700377.jpg', '#2e8b57', 'Hội An'),
        ('Bánh Mì', 'Bánh mì là món ăn đường phố nổi tiếng của Việt Nam, với nhiều loại nhân khác nhau.', 'Bánh mì, Thịt, Rau sống, Nước chấm', '15.000 - 25.000 VNĐ', 'https://cdn.xanhsm.com/2025/01/125f9835-banh-mi-sai-gon-thumb.jpg', '#ffa500', 'Việt Nam'),
        ('Gỏi Cuốn', 'Gỏi cuốn là món ăn nhẹ, bao gồm tôm, thịt, rau sống và bún, cuốn trong bánh tráng.', 'Tôm, Thịt, Rau sống, Bún, Bánh tráng', '20.000 - 30.000 VNĐ', 'https://cdn2.fptshop.com.vn/unsafe/1920x0/filters:format(webp):quality(75)/2023_10_23_638336957766719361_cach-lam-goi-cuon.jpg', '#8b4513', 'Miền Nam'),
        ('Bánh Xèo', 'Bánh xèo là món ăn đặc sản miền Trung, với lớp bánh mỏng, giòn, nhân tôm, thịt và giá đỗ.', 'Bánh xèo, Tôm, Thịt, Giá đỗ, Rau sống', '25.000 - 35.000 VNĐ', 'https://i-giadinh.vnecdn.net/2023/09/19/Buoc-10-Thanh-pham-1-1-5225-1695107554.jpg', '#ffd700', 'Miền Trung'),
        ('Phở', 'Phở là món ăn truyền thống của Việt Nam, với nước dùng thơm ngon, bánh phở mềm và thịt bò hoặc gà.', 'Bánh phở, Thịt bò, Nước dùng, Rau thơm', '30.000 - 50.000 VNĐ', 'https://giavichinsu.com/wp-content/uploads/2024/01/cach-nau-pho-bo.jpg', '#ff6347', 'Miền Bắc'),
        ('Bún Bò Huế', 'Bún bò Huế là món ăn đặc sản của Huế, với nước dùng đậm đà, sợi bún to và thịt bò mềm.', 'Bún, Thịt bò, Giò heo, Hành tây, Rau sống', '40.000 - 60.000 VNĐ', 'https://file.hstatic.net/200000700229/article/bun-bo-hue-1_da318989e7c2493f9e2c3e010e722466.jpg', '#8a2be2', 'Huế'),
        ('Bánh Canh', 'Bánh canh là món ăn đặc sản của miền Trung, với sợi bánh dày, nước dùng đậm đà và các loại topping phong phú.', 'Bánh canh, Tôm, Cua, Nghêu, Rau sống', '35.000 - 55.000 VNĐ', 'https://www.huongnghiepaau.com/wp-content/uploads/2018/01/banh-canh-cua.jpg', '#ff4500', 'Miền Trung');
        `
		_, err := db.Exec(insertSQL)
		if err != nil {
			log.Fatalf("Không thể chèn dữ liệu: %v", err)
		}
		log.Println("Chèn dữ liệu thành công!")
	} else {
		log.Println("Bảng đã có dữ liệu, bỏ qua bước chèn.")
	}
}

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
	migrateAndSeed(db)

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
	// ... (phần code này giữ nguyên)
	rows, err := db.Query("SELECT id, name, description, ingredients, price, image_url, color, region FROM foods ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	foods := []Food{}
	for rows.Next() {
		var f Food
		if err := rows.Scan(&f.ID, &f.Name, &f.Description, &f.Ingredients, &f.Price, &f.ImageURL, &f.Color, &f.Region); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		foods = append(foods, f)
	}
	c.JSON(http.StatusOK, foods)
}
