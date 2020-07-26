package config

type Config struct{
	DB *DBConfig
}

type DBConfig struct{
	Dialect string
	Host string
	Port string
	Username string
	Password string
	Name string
	Charset string
}

func GetConfig() *Config{
	return &Config{
		DB: &DBConfig{
			Dialect :"",
			Host : "",
			Port:"",
			Username:"",
			Password:"",
			Name:"",
			Charset:"",
		},
	}
}