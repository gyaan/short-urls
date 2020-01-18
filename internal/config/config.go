package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

var (
	conf       *Config
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "../../")
)

//Config
type Config struct {
	ApplicationPort           string `json:"application_port" mapstructure:"application_port"`
	MongoDbConnectionUrl      string `json:"mongo_db_connection_url" mapstructure:"mongo_db_connection_url"`
	MongoDatabaseName         string `json:"mongo_database_name" mapstructure:"mongo_database_name"`
	ShortUrlExpiryTime        int    `json:"short_url_expiry_time" mapstructure:"short_url_expiry_time"` // time in minutes
	BaseUrl                   string `json:"base_url" mapstructure:"base_url"`
	MinimumShortUrlIdentifier int    `json:"minimum_short_url_identifier" mapstructure:"minimum_short_url_identifier"`
	MongoContextTimeout       int    `json:"mongo_context_timeout" mapstructure:"mongo_context_timeout"`
	JWTSecret                 string `json:"jwt_secret" mapstructure:"jwt_secret"`
	TokenExpiryTime           int64  `json:"token_expiry_time" mapstructure:"token_expiry_time"`
}

// NewConfig returns creates new application config
func New() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(basePath + "/config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return conf
}

// GetConf returns config
func GetConf() *Config {
	return conf
}
