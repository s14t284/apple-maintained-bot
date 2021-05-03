package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/domain/model"

	"github.com/labstack/gommon/log"
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

func initializeTable(migrator gorm.Migrator, tableInterface interface{}) error {
	if !migrator.HasTable(tableInterface) {
		err := migrator.CreateTable(tableInterface)
		return err
	}
	return nil
}

func initializeTables(dbClient *gorm.DB) error {
	errStr := ""
	migrator := dbClient.Migrator()
	err := initializeTable(migrator, &model.Mac{})
	if err != nil {
		errStr += "initialize mac table error: " + err.Error() + "\n"
	}
	err = initializeTable(migrator, &model.IPad{})
	if err != nil {
		errStr += "initialize ipad table error: " + err.Error() + "\n"
	}
	err = initializeTable(migrator, &model.Watch{})
	if err != nil {
		errStr += "initialize watch table error: " + err.Error() + "\n"
	}

	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
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

	err = initializeTables(dbClient)
	if err != nil {
		return nil, err
	}
	return &SQLClient{Client: dbClient}, nil
}
