package postgress

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type Config struct {
	Host string
	Port string
	User string
	Password string
	Name string
}


func LoadConfigFromEnv() *Config{
	return &Config{
		Host: getenv("db_host","localhost"),
		Port: getenv("db_port", "5433"),
		User: getenv("db_user","app_user"),
		Password: getenv("db_password", "app_password"),
		Name: getenv("db_name", "todo_list"),
	}
}

func (cfg *Config) ConnectFromEnv() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	return db, nil
}

func getenv(key, df string) string {
	v := os.Getenv(key)
	if v == "" {
		return df
	}
	return v
}