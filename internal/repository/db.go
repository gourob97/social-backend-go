package repository

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)	

func ConnectDB(url string) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}