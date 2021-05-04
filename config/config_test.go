package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	testPsqlPort        = 3306
	testPsqlHost        = "localhost"
	testUserName        = "user"
	testPassword        = "password"
	testDatabase        = "db"
	testSlackChannel    = "#geneeral"
	testSlackUserName   = "slackUser"
	testSlackIcon       = ":+1:"
	testSlackWebHookURL = "https://hooks.slack.com/services/dummy"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestAllSuites(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (cts *ConfigTestSuite) BeforeTest(suiteName string, testName string) {
	_ = os.Setenv("PSQL_PORT", strconv.Itoa(testPsqlPort))
	_ = os.Setenv("PSQL_HOST", testPsqlHost)
	_ = os.Setenv("PSQL_USER_NAME", testUserName)
	_ = os.Setenv("PSQL_PASSWORD", testPassword)
	_ = os.Setenv("PSQL_DATABASE", testDatabase)
	_ = os.Setenv("SLACK_CHANNEL", testSlackChannel)
	_ = os.Setenv("SLACK_USER_NAME", testSlackUserName)
	_ = os.Setenv("SLACK_ICON", testSlackIcon)
	_ = os.Setenv("SLACK_WEBHOOK_URL", testSlackWebHookURL)
}

func (cts *ConfigTestSuite) TestLoadConfig() {
	a := assert.New(cts.T())
	lc, err := LoadConfig()
	if err != nil {
		cts.T().FailNow()
	}
	a.Equal(testPsqlPort, lc.PsqlConfig.Port)
	a.Equal(testPsqlHost, lc.PsqlConfig.Host)
	a.Equal(testUserName, lc.PsqlConfig.UserName)
	a.Equal(testPassword, lc.PsqlConfig.Password)
	a.Equal(testDatabase, lc.PsqlConfig.Database)
	a.Equal(testSlackChannel, lc.SlackNotifyConfig.Channel)
	a.Equal(testSlackUserName, lc.SlackNotifyConfig.UserName)
	a.Equal(testSlackIcon, lc.SlackNotifyConfig.IconEmoji)
	a.Equal(testSlackWebHookURL, lc.SlackNotifyConfig.WebHookURL)
}

func (cts *ConfigTestSuite) TestCreatePsqlConfigWhenAllEmpty() {
	a := assert.New(cts.T())
	_ = os.Setenv("PSQL_PORT", "")
	_ = os.Setenv("PSQL_HOST", "")
	_ = os.Setenv("PSQL_USER_NAME", "")
	_ = os.Setenv("PSQL_PASSWORD", "")
	_ = os.Setenv("PSQL_DATABASE", "")
	pc, err := createPsqlConfig()
	if err != nil {
		cts.T().FailNow()
	}

	a.Equal(5432, pc.Port)
	a.Equal("", pc.Host)
	a.Equal("", pc.UserName)
	a.Equal("", pc.Password)
	a.Equal("", pc.Database)
}

func (cts *ConfigTestSuite) TestCreateSlackNotifyConfigWhenAllEmpty() {
	a := assert.New(cts.T())
	_ = os.Setenv("SLACK_CHANNEL", "")
	_ = os.Setenv("SLACK_USER_NAME", "")
	_ = os.Setenv("SLACK_ICON", "")
	_ = os.Setenv("SLACK_WEBHOOK_URL", "https://dummy.com")
	snc, err := createSlackNotifyConfig()
	if err != nil {
		cts.T().FailNow()
	}

	a.Equal("#random", snc.Channel)
	a.Equal("AppleMaintainedBot", snc.UserName)
	a.Equal(":apple:", snc.IconEmoji)
	a.Equal("https://dummy.com", snc.WebHookURL)
}

func (cts *ConfigTestSuite) TestCreateSlackNotifyConfigWhenParamsImperfects() {
	a := assert.New(cts.T())
	channel := "general"
	icon := "icon"
	_ = os.Setenv("SLACK_CHANNEL", channel)
	_ = os.Setenv("SLACK_ICON", icon)
	_ = os.Setenv("SLACK_WEBHOOK_URL", "https://dummy.com")
	snc, err := createSlackNotifyConfig()
	if err != nil {
		cts.T().FailNow()
	}

	a.Equal("#"+channel, snc.Channel)
	a.Equal(":"+icon+":", snc.IconEmoji)
}
