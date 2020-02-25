package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gorm.NowFunc = func() time.Time {
		var location, err = time.LoadLocation("Asia/Jakarta")
		if err != nil {
			return time.Now().UTC()
		}
		return time.Now().In(location)
	}
}

type DatabaseConfig struct {
	Driver            string
	Host              string
	Port              int
	User              string
	Password          string
	DatabaseName      string
	MaxIdleConnection int
	MaxOpenConnection int
}

type Config struct {
	Database      DatabaseConfig
	IsDevelopment bool
}

func New(db DatabaseConfig, isDev bool) *Config {
	return &Config{
		Database:      db,
		IsDevelopment: isDev,
	}
}

func (conf Config) DB() (*gorm.DB, error) {
	if conf.Database.Driver == "postgres" {
		args := fmt.Sprintf(
			"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.DatabaseName,
		)
		db, err := gorm.Open(conf.Database.Driver, args)
		if err != nil {
			return db, err
		}
		db.DB().SetMaxIdleConns(conf.Database.MaxIdleConnection)
		db.DB().SetMaxOpenConns(conf.Database.MaxOpenConnection)
		db.DB().SetConnMaxLifetime(time.Hour)
		db.LogMode(conf.IsDevelopment)
		return db, nil
	}
	return nil, fmt.Errorf("database driver %s not supported", conf.Database.Driver)
}

func NewPostgreSQLConfig() DatabaseConfig {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Error loading DB_PORT")
	}

	return DatabaseConfig{
		Driver:            os.Getenv("DB_CONNECTION"),
		Host:              os.Getenv("DB_HOST"),
		Port:              dbPort,
		User:              os.Getenv("DB_USERNAME"),
		Password:          os.Getenv("DB_PASSWORD"),
		DatabaseName:      os.Getenv("DB_DATABASE"),
		MaxIdleConnection: 10,
		MaxOpenConnection: 100,
	}
}
