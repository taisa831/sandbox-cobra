package config

var Conf Config

type Config struct {
	DBConfig dbConfig
}

type dbConfig struct {
	Schema string `env:"SCHEMA" envDefault:"sandbox-cobra"`
	User string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:"password"`
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"3306"`
}