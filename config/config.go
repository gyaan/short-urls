package config

var conf *Config

type Config struct {
	ApplicationPort      string `json:"application_port"`
	MongoDbConnectionUrl string `json:"mongo_db_connection_url"`
	MongoDatabaseName    string `json:"mongo_database_name"`
	ShortUrlExpiryTime   int    `json:"short_url_expiry_time"` // time in hours
}

// NewConfig returns creates new application config
func NewConfig(config Config) *Config {

	conf = &Config{
		ApplicationPort:      config.ApplicationPort,
		MongoDbConnectionUrl: config.MongoDbConnectionUrl,
		MongoDatabaseName:    config.MongoDatabaseName,
		ShortUrlExpiryTime:   config.ShortUrlExpiryTime,
	}
	return conf
}

// GetConf returns config
func GetConf() *Config {
	return conf
}
