package config

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() {
  dsn := "root:root@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  DB = db
}

func GetDB() *gorm.DB {
  return DB
}
