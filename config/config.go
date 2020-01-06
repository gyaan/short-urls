package config

var conf *Config

type Config struct {
	ApplicationPort      string `json:"application_port"`
	MongoDbConnectionUrl string `json:"mongo_db_connection_url"`
}

// NewConfig returns creates new application config
func NewConfig(config Config) *Config {

	conf = &Config{
		ApplicationPort:      config.ApplicationPort,
		MongoDbConnectionUrl: config.MongoDbConnectionUrl,
	}
	return conf
}

// GetConf returns config
func GetConf() *Config {
	return conf
}
