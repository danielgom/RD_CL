package config

import (
	"bytes"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/viper"
)

// Config is where the global configuration is stored.
type Config struct {
	DB         DB         `mapstructure:"DB"`
	SMTP       SMTP       `mapstructure:"SMTP"`
	Logger     Logger     `mapstructure:"Logger"`
	Server     Server     `mapstructure:"Server"`
	JWT        JWT        `mapstructure:"JWT"`
	FileSystem FileSystem `mapstructure:"FileSystem"`
}

type DB struct {
	Name     string `mapstructure:"Name"`
	Port     string `mapstructure:"Port"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Host     string `mapstructure:"Host"`
}

type SMTP struct {
	From     string `mapstructure:"FROM_MAIL"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type Server struct {
	Environment string `mapstructure:"Environment"`
	Address     string `mapstructure:"Address"`
	Timeout     string `mapstructure:"Timeout"`
}

type Logger struct {
	Level string `mapstructure:"Level"`
}

type JWT struct {
	Expiration int32  `mapstructure:"Expiration"`
	Key        string `mapstructure:"Key"`
}

type FileSystem struct {
	BaseDir string `mapstructure:"BaseDir"`
}

// Load gets the configuration in from .env files and stores the in Config struct.
func Load() *Config {
	var c Config
	var once sync.Once

	once.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			panic("cannot get current working directory")
		}

		viper.AddConfigPath(dir + "/internal/config/cfg")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalln("Could not read configuration:", err.Error())
		}

		err = viper.Unmarshal(&c)
		if err != nil {
			log.Fatalln("Could not unmarshal to Config struct:", err.Error())
		}

		c.setBaseDir()
	})

	return &c
}

func (c *Config) setBaseDir() {
	if c.FileSystem.BaseDir != "" {
		return
	}

	if c.IsDevelopment() || c.IsCI() {
		log.Println("Getting BaseDir from git")
		cmd := exec.Command("git", "rev-parse", "--show-toplevel")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout
		if err := cmd.Run(); err != nil {
			log.Fatalln("Could not run git:", err.Error())
		} else {
			root := bytes.TrimRight(stdout.Bytes(), "\n")
			c.FileSystem.BaseDir = string(root)
		}
	} else {
		log.Println("Getting BaseDir from cwd")
		if cwd, err := os.Getwd(); err == nil {
			c.FileSystem.BaseDir = cwd
		} else {
			log.Fatalln("Failed to get current directory: ", err.Error())
		}
	}
}

var dev, ci = "development", "ci"

// IsDevelopment returns true when running in a local development environment.
func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == dev
}

// IsCI returns true when running in GitLab CI environment.
func (c *Config) IsCI() bool {
	return c.Server.Environment == ci
}

func (c *Config) ServerTimeout() time.Duration {
	if c.Server.Timeout == "" {
		return time.Duration(0)
	}
	d, err := time.ParseDuration(c.Server.Timeout)
	if err != nil {
		panic("Invalid Server.Timeout: " + err.Error())
	}
	return d
}

func (c *Config) ServerAddress() string {
	if c.Server.Address == "" {
		return "localhost:8080"
	}
	return c.Server.Address
}

func (c *Config) LogLevel() slog.Level {
	level := c.Logger.Level
	switch level {
	case "", "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic("Invalid Log.Level: " + level)
	}
}
