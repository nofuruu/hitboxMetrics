package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"go-test/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UploadController struct {
	DB *gorm.DB
}

func (u *UploadController) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve video file"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".mp4" && ext != ".mkv" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type. Only .mp4 and .mkv are allowed."})
		return
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)
	savePath := filepath.Join("uploads", newFilename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video file"})
		return
	}

	match := models.Match{
		OriginalName: file.Filename,
		SavedName:    newFilename,
		FilePath:     savePath,
		Status:       "pending",
	}

	if err := u.DB.Create(&match).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mencatat ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Video berhasil diunggah dan masuk antrean",
		"match_id": match.ID,
		"status":   match.Status,
	})
}
