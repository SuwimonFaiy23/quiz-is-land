package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var DB DBConfig
var Conn *gorm.DB

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ไม่สามารถอ่าน config: %v", err)
	}

	DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Name:     viper.GetString("database.name"),
	}
}

func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB.User, DB.Password, DB.Host, DB.Port, DB.Name)
}

func ConnectDatabase() *gorm.DB {
	dsn := GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("เชื่อมต่อ DB ไม่ได้:", err)
	}

	return db
}
