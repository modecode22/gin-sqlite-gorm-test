package database

import (
    "gin-sqlite-gorm-test/internal/models"

    "gorm.io/gorm"
    "github.com/glebarez/sqlite" // Change this import
)

var DB *gorm.DB

func InitDB() error {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) // Use sqlite.Open here
    if err != nil {
        return err
    }

    // Auto migrate the schema
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        return err
    }

    return nil
}