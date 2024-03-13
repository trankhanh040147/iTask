package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	App       AppConfig
	Mysql     MysqlConfig
	AWS       AWSConfig
	Redis     RedisConfig
	Google    GoogleConfig
	Email     EmailSenderConfig
	CronSpec  CronSpec
	GoogleMap GoogleMap
	Momo      Momo
}

type CronSpec struct {
	UpdateStatusBooking string
}

type EmailSenderConfig struct {
	EmailSenderName     string
	EmailSenderAddress  string
	EmailSenderPassword string
}

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
	Endpoint     any
}

type AppConfig struct {
	Version      string
	Port         string
	Mode         string
	Secret       string
	MigrationURL string
}
type MysqlConfig struct {
	Host            string
	ContainerName   string
	Port            string
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime string
}

type AWSConfig struct {
	Region         string
	APIKey         string
	SecretKey      string
	S3Bucket       string
	S3Domain       string
	S3FolderImages string
}
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type GoogleMap struct {
	APIKey  string
	BaseURL string
}

type Momo struct {
	EndPoint    string
	AccessKey   string
	SecretKey   string
	PartnerCode string
	RedirectURL string
	IpURL       string
	Lang        string
	PartnerName string
	StoreId     string
	AutoCapture bool
	RequestType string
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// check is not found file config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config // Unmarshal data config have get in file config then get into c
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// trigger cicd
