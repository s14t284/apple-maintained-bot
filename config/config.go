package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// PsqlConfig psqlに接続するための設定値
type PsqlConfig struct {
	Host     string
	UserName string
	Password string
	Port     int
	Database string
}

// SlackNotifyConfig slackに通知を送るための設定値
type SlackNotifyConfig struct {
	Channel    string
	UserName   string
	IconEmoji  string
	WebHookURL string
}

// Config 設定値をまとめた構造体
type Config struct {
	PsqlConfig        PsqlConfig
	SlackNotifyConfig SlackNotifyConfig
}

func createPsqlConfig() (*PsqlConfig, error) {
	port, err := strconv.Atoi(os.Getenv("PSQL_PORT"))
	if err != nil {
		log.Errorf("can't load port of psql host")
		port = 5432 // set psql default port
	}
	psqlConfig := &PsqlConfig{
		Host:     os.Getenv("PSQL_HOST"),
		UserName: os.Getenv("PSQL_USER_NAME"),
		Password: os.Getenv("PSQL_PASSWORD"),
		Port:     port,
		Database: os.Getenv("PSQL_DATABASE"),
	}
	return psqlConfig, nil
}

func createSlackNotifyConfig() (*SlackNotifyConfig, error) {
	channel := os.Getenv("SLACK_CHANNEL")
	if channel == "" {
		channel = "#random"
	} else if !strings.HasPrefix(channel, "#") {
		channel = "#" + channel
	}

	userName := os.Getenv("SLACK_USER_NAME")
	if userName == "" {
		userName = "AppleMaintainedBot"
	}

	iconEmoji := os.Getenv("SLACK_ICON")
	if iconEmoji == "" {
		iconEmoji = ":apple:"
	}
	if !strings.HasPrefix(iconEmoji, ":") {
		iconEmoji = ":" + iconEmoji
	}
	if !strings.HasSuffix(iconEmoji, ":") {
		iconEmoji = iconEmoji + ":"
	}

	webHookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	if webHookUrl == "" {
		return nil, fmt.Errorf("failed to load SLACK_WEBHOOK_URL")
	}

	return &SlackNotifyConfig{
		Channel:    channel,
		UserName:   userName,
		IconEmoji:  iconEmoji,
		WebHookURL: webHookUrl,
	}, nil
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

	// slack
	slackConfig, err := createSlackNotifyConfig()
	if err != nil {
		return nil, err
	}

	// timezone
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)

	return &Config{
		PsqlConfig:        *psqlConfig,
		SlackNotifyConfig: *slackConfig,
	}, nil
}
