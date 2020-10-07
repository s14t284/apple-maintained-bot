package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
)

// Config 設定値をまとめた構造体
type Config struct {
	PsqlConfig infrastructure.PsqlConfig
}

func createPsqlConfig() (infrastructure.PsqlConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Errorf("can't load port of psql host")
		port = 5432 // set psql default port
	}
	psqlConfig := infrastructure.PsqlConfig{
		Host:     os.Getenv("HOST"),
		UserName: os.Getenv("USER_NAME"),
		Password: os.Getenv("PASSWORD"),
		Port:     port,
		Database: os.Getenv("DATABASE"),
	}
	return psqlConfig, nil
}

// LoadConfig 設定値を読み込む
func LoadConfig() (*Config, error) {
	// ローカルで実行した場合、.envから環境変数を読み込む
	err := godotenv.Load(".env")
	if err == nil {
		log.Infof("load environment variables from `.env` file")
	}

	// psql
	psqlConfig, err := createPsqlConfig()
	if err != nil {
		return nil, err
	}

	// timezone
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	fmt.Println(time.Now())

	return &Config{
		PsqlConfig: psqlConfig,
	}, nil
}
