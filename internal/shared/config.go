package shared

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password,omitempty"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password,omitempty"` // Optional, can be empty
	} `yaml:"redis"`
	JWTSecret string `yaml:"jwt_secret"`
}

var Acfg *Config

func init() {
	var err error
	Acfg, err = loadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}
}

func loadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
func FiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber API",
		ErrorHandler:  ErrorHandler,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	}
}
