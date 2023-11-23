package config

type AuthConfig struct {
	PublicKeyPath  string `env:"PUBLIC_KEY_PATH" envDefault:""`
	PrivateKeyPath string `env:"PRIVATE_KEY_PATH" envDefault:""`
}