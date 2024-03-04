package seed

import (
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"

	"gorm.io/gorm"

	"time"
)

var users = []models.User{
	{
		Username:  "cemongdjtnk",
		Email:     "cemongdjatniko@mail.com",
		Password:  "cemong123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Username:  "frankiedjtnk",
		Email:     "frankiedjatniko@mail.com",
		Password:  "frankie123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var photos = []models.Photo{
	{
		Title:     "My favorit books",
		Caption:   "This is my favorit books lately that really motivated me",
		PhotoURL:  "https://drive.google.com/file/d/1gZBPqKtEMraB-8o2uFZxwhY8G9zUc4m_/view?usp=sharing",
		UserID:    3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Title:     "Beutiful Scenary",
		Caption:   "This image I capture in Breksi Yogyakarta",
		PhotoURL:  "https://drive.google.com/file/d/1Lon6ceNtrHvJPwCreMy2S1Z5qxMLjfCz/view?usp=sharing",
		UserID:    3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Title:     "Chinese Porridge with pek cam kee",
		Caption:   "Nailed it to made this comfort food!",
		PhotoURL:  "https://drive.google.com/drive/my-drive?hl=id",
		UserID:    4,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func Load(db *gorm.DB) {
	db.Create(&users)
	db.Create(&photos)
}
