package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(postgres.Open("postgres://postgres:keeplowkey@localhost:5432/go_task_btpn?sslmode=disable"))
    if err != nil {
        panic("failed to connect database")
    }

    err = database.AutoMigrate(&User{})
    if err != nil {
        panic("failed to migrate User table schema")
    }

    err = database.AutoMigrate(&Photo{})
    if err != nil {
        panic("failed to migrate Photo table schema")
    }

    DB = database
}
