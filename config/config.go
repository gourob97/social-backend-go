package config

type AppConfig struct {
	Port      string
	JWTSecret string
}

type DbConfig struct {
	Host string
	Port string
	User string
	Pass string
	Schema string
}

type Config struct {
	App *AppConfig
	Db  *DbConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func Db() *DbConfig {
	return config.Db
}


func LoadConfig() *Config {
	setDefaultConfig()
	
	return &config
}


func setDefaultConfig() {
	config = Config{
		App: &AppConfig{
			Port:      "8080",
			JWTSecret: "defaultsecret",
		},
		Db: &DbConfig{
			Host:   "localhost",
			Port:   "3306",
			User:   "socialuser",
			Pass:   "socialpassword",
			Schema: "social",
		},
	}
}