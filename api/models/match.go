package models

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	OriginalName string `json:"originalName"`
	SavedName    string `json:"savedName"`
	FilePath     string `json:"filePath"`
	Status       string `json:"status"`
}
