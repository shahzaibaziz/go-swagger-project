package mysql

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/shahzaibaziz/GOProject/config"
	"github.com/shahzaibaziz/GOProject/db"
)

func init() {
	db.Register("mysql", NewClient)
}

// client struct for mysql client
type client struct {
	conn gorm.DB
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString(config.DBUser), viper.GetString(config.DBPassword), viper.GetString(config.DBHost), viper.GetString(config.DBPort), viper.GetString(config.DBName))

	conn, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("database Connection error %s", databaseURL)
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	if err = conn.AutoMigrate(); err != nil {
		return nil, errors.Wrap(err, "failed to create tables")
	}

	return &client{conn: *conn}, nil
}
