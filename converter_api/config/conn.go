package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	if err := os.Getenv("MYSQL_HOST"); err == "" {
		_ = fmt.Errorf("the SECRET_TOKEN not found on environment variables")
		return nil, fmt.Errorf("the MYSQL_HOST not found on environment variables")
	}
	var host = os.Getenv("MYSQL_HOST")
	var port = os.Getenv("MYSQL_PORT")
	var user = os.Getenv("MYSQL_USER")
	var dbname = os.Getenv("MYSQL_DATABASE")
	var password = os.Getenv("MYSQL_PASSWORD")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
