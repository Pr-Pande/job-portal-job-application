package config

type AppConfig struct {
	Port string `env:"APP_PORT,required=true"`
}
