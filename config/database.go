package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEvn := godotenv.Load()
	if errEvn != nil {
		panic("Faile to load env file env")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to create to Connect Database")
	}
	return db
}
func closeConnect(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Faile to close database")
	}
	dbSql.Close()
}
