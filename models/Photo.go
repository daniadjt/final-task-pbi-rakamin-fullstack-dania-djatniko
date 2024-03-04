package models

import (
    "time"
)

type Photo struct {
    ID        int64     `gorm:"primaryKey" json:"id"`
    Title     string    `gorm:"type:varchar(300)" json:"title"`
    Caption   string    `gorm:"type:varchar(300)" json:"caption"`
    PhotoURL  string    `gorm:"type:varchar(300)" json:"photo_url"`
    UserID    uint     `json:"user_id"`
    User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}