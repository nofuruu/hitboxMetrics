package main

import (
	"api/models"
	"log"
	"time"

	"hitboxmetrics/api/controllers"
	"hitboxmetrics/api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("hitbox.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.Match{})

    uploadCtrl := &controllers.UploadController{DB: db}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hitboxMetrics API is online!"})
		})

        api.POST("/upload", uploadCtrl.HandleUpload)
	}

    log.Println("Server berjalan di http://localhost:8000")
	r.Run(":8080")
}
