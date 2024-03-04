package models

import (
    "time"
)

type User struct {
    UserID    int64     `gorm:"primaryKey" json:"user_id"`
    Username  string    `gorm:"required" json:"username"`
    Email     string    `gorm:"type:varchar(300);unique;required" json:"email"`
    Password  string    `gorm:"type:varchar(300);required;min:6" json:"password"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}