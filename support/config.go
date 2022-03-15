package support

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	file        string
	Env         string
	App         string
	Version     string
	DB_POSTGRES string
	DB_MYSQL    string
	ServerPort  string
	Secret      string
}

func LoadConfig() *Config {
	return loadByFile(".env")
}

func loadByFile(envFile string) *Config {
	config := Config{file: ".env"}
	err := config.loadEnvs(envFile)
	if err != nil {
		panic(err)
	}
	config.setEnvs()

	return &config
}

func (c *Config) SetIsTesting() {
	c.file = ".env"
}

func (c *Config) loadEnvs(envFile string) error {
	return godotenv.Load(envFile)
}

func (c *Config) setEnvs() {
	c.Env = os.Getenv("ENVIRONMENT")
	c.App = os.Getenv("APP")
	c.Version = os.Getenv("VERSION")
	c.DB_POSTGRES = os.Getenv("DB_POSTGRES")
	c.DB_MYSQL = os.Getenv("DB_MYSQL")
	c.ServerPort = os.Getenv("SERVER_PORT")
	c.Secret = os.Getenv("SECRET")
}
