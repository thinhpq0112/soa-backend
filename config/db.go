package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func NewConn() (*gorm.DB, error) {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Database environment variables are missing or empty")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	sqlDb, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
