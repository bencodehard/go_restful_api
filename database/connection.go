package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	viper.AutomaticEnv()
	// viper.SetConfigFile(".env")
	// viper.AddConfigPath("./.env")

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Access variables
	dbHost := viper.GetString("MARIA_DB_HOST")
	dbPort := viper.GetInt("MARIA_DB_PORT")
	dbUsername := viper.GetString("MARIA_DB_USER")
	dbPassword := viper.GetString("MARIA_DB_PASSWORD")
	dbName := viper.GetString("MARIA_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect Db successful")
	}
	db = d

}
