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
	MongoDatabaseUserName     string `json:"mongo_database_user_name" mapstructure:"mongo_database_username"`
	MongoDatabasePassword     string `json:"mongo_database_password" mapstructure:"mongo_database_password"`
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

	//set mongo env variable default config values will overwrite with env variable
	err = viper.BindEnv("mongo_database_name", "MONGO_INITDB_DATABASE")
	if err != nil {
		log.Printf("not able to set env variable for mongo_database_name %v", err)
	}

	err = viper.BindEnv("mongo_database_username", "MONGO_INITDB_ROOT_USERNAME")
	if err != nil {
		log.Printf("not able to set env variable for mongo_database_username %v", err)
	}

	err = viper.BindEnv("mongo_database_password", "MONGO_INITDB_ROOT_PASSWORD")
	if err != nil {
		log.Printf("not able to set env variable for mongo_database_password %v", err)
	}

	return conf
}

// GetConf returns config
func GetConf() *Config {
	return conf
}
