package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"yxb.com/gin-scaffold/v2/pkg/logger"
)

var C = new(Config)

func LoadConf(f string) {
	viper.SetConfigFile(f)
	if err := viper.ReadInConfig(); err != nil {
		logger.WithContext(context.Background()).Fatalf("read config file failed. err: %v", err.Error())
	}
	//viper.UnmarshalKey("mysql", &C.MySQL)
	if err := viper.Unmarshal(C); err != nil {
		logger.WithContext(context.Background()).Fatalf("unmarshal data failed. err: %v", err.Error())
	}
}

func PrintWithJson() {
	b, err := json.MarshalIndent(C, "", " ")
	if err != nil {
		os.Stdout.WriteString("[Config] Json marshal error " + err.Error())
		return
	}
	os.Stdout.WriteString(string(b) + "\n")

}

type Config struct {
	HttpServer HttpServer
	MySQL      MySQL
	Gorm       Gorm
	Redis      Redis
	Log        Log
	Casbin     Casbin
	Jwt        Jwt
	Sqlite3    Sqlite3
}

type HttpServer struct {
	Host string
	Port int
}

type Casbin struct {
	Enabled bool
	Model   string
}

type LogHook string

func (h LogHook) IsGorm() bool {
	return h == "gorm"
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
	RotationCount int
	RotationTime  int
}

type Gorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type MySQL struct {
	Host       string
	Port       int
	User       string
	Pwd        string
	Dbname     string
	Parameters string
}

func (m MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		m.User, m.Pwd, m.Host, m.Port, m.Dbname, m.Parameters)
}

type Sqlite3 struct {
	Path string
}

func (a Sqlite3) DSN() string {
	return a.Path
}

type Redis struct {
	Host string
	Port int
}

type Jwt struct {
	Enabled       int
	SigningMethod string
	SigningKey    string
}
