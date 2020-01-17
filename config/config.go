package config

var conf *Config

type Config struct {
	ApplicationPort           string `json:"application_port"`
	MongoDbConnectionUrl      string `json:"mongo_db_connection_url"`
	MongoDatabaseName         string `json:"mongo_database_name"`
	ShortUrlExpiryTime        int    `json:"short_url_expiry_time"` // time in hours
	BaseUrl                   string `json:"base_url"`
	MinimumShortUrlIdentifier int    `json:"minimum_short_url_identifier"`
	MongoContextTimeout       int    `json:"mongo_context_timeout"`
	JWTSecret                 string `json:"jwt_secret"`
	TokenExpiryTime           int64  `json:"token_expiry_time"`
}

// NewConfig returns creates new application config
func NewConfig(config Config) *Config {

	conf = &Config{
		ApplicationPort:           config.ApplicationPort,
		MongoDbConnectionUrl:      config.MongoDbConnectionUrl,
		MongoDatabaseName:         config.MongoDatabaseName,
		ShortUrlExpiryTime:        config.ShortUrlExpiryTime,
		BaseUrl:                   config.BaseUrl,
		MinimumShortUrlIdentifier: config.MinimumShortUrlIdentifier,
		MongoContextTimeout:       config.MongoContextTimeout,
		JWTSecret:                 config.JWTSecret,
		TokenExpiryTime:           config.TokenExpiryTime,
	}
	return conf
}

// GetConf returns config
func GetConf() *Config {
	return conf
}
