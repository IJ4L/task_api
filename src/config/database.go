package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:@tcp(127.0.0.1:3306)/taskManager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}
