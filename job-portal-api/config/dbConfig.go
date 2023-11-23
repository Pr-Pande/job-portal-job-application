package config

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST,required=true"`
	UserName string `env:"POSTGRES_USER,required=true"`
	Password string `env:"POSTGRES_PASSWORD,required=true"`
	DBName   string `env:"POSTGRES_DBNAME,required=true"`
	Port     string `env:"POSTGRES_PORT,default=5432"`
	SSLMode  string `env:"POSTGRES_SSL_MODE,default=false"`
	TimeZone     string `env:"POSTGRES_TIMEZONE,default=Asia/Shanghai"`
}