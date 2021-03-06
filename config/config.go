package config

type Config struct{
	DB *DBConfig
}

type DBConfig struct{
	Dialect string
	Host string
	Port int
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
			Port:0,
			Username:"",
			Password:"",
			Name:"",
			Charset:"",
		},
	}
}