package database

import (
	"fmt"
	"os"

	"github.com/s14t284/apple-maitained-bot/config"

	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SQLClient psqlの接続クライアント
type SQLClient struct {
	Client *gorm.DB
}

func createDataSourceName(config config.PsqlConfig) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=require TimeZone=Asia/Tokyo",
		config.UserName, config.Password, config.Host, config.Port, config.Database)
}

func initializeTable(migrator gorm.Migrator, tableName string, tableInterface interface{}) {
	if !migrator.HasTable(tableInterface) {
		migrator.CreateTable(tableInterface)
	}
}

func initializeTables(dbClient *gorm.DB) {
	migrator := dbClient.Migrator()
	initializeTable(migrator, "mac", &model.Mac{})
	initializeTable(migrator, "ipad", &model.IPad{})
	initializeTable(migrator, "watch", &model.Watch{})
}

// PsqlNewClientImpl psqlに接続したgormクライアントを返却
func PsqlNewClientImpl(c config.PsqlConfig) (*SQLClient, error) {
	// deploy先では環境変数から接続先情報を取得できる
	dataSourceName := os.Getenv("DATABASE_URL")
	// localでは予め .env に入力しておいた情報から接続先を特定
	if dataSourceName == "" {
		dataSourceName = createDataSourceName(c)
	}
	dbClient, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	initializeTables(dbClient)
	return &SQLClient{Client: dbClient}, nil
}
