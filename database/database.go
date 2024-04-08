package database

import (
	"fmt"
	"meeting3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/bootcamp_ramadhan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("========================= Connection Error !! ======================= \n", err)
	}

	DB = db
	fmt.Println("========================= Connection Success ... ======================= ")

	db.AutoMigrate(models.Admin{}, models.User{})
  }


