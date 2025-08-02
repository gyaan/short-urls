package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	conf       *Config
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Join(filepath.Dir(b), "../../")
)

// Config holds application configuration settings
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

// New creates and returns a new application configuration
func New() *Config {
	v := viper.New()

	// Bind environment variables for MongoDB configuration
	if err := v.BindEnv("mongo_database_name", "MONGO_INITDB_DATABASE"); err != nil {
		log.Printf("Failed to bind env variable for mongo_database_name: %v", err)
	}

	if err := v.BindEnv("mongo_database_username", "MONGO_INITDB_ROOT_USERNAME"); err != nil {
		log.Printf("Failed to bind env variable for mongo_database_username: %v", err)
	}

	if err := v.BindEnv("mongo_database_password", "MONGO_INITDB_ROOT_PASSWORD"); err != nil {
		log.Printf("Failed to bind env variable for mongo_database_password: %v", err)
	}

	if err := v.BindEnv("mongo_db_connection_url", "MONGO_HOST"); err != nil {
		log.Printf("Failed to bind env variable for mongo_db_connection_url: %v", err)
	}

	// Configure viper to read YAML config file
	v.SetConfigType("yaml")
	v.AddConfigPath(basePath + "/config")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	log.Printf("Configuration loaded: %+v", conf)
	return conf
}

// GetConf returns the current configuration instance
func GetConf() *Config {
	return conf
}
