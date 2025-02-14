package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/spf13/viper"
)

type Config struct {
	MongoURI      string `mapstructure:"MONGODB_URI"`
	CloudinaryURL string `mapstructure:"CLOUDINARY_URL"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

func InitCloudinary(cfg Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(cfg.CloudinaryURL)
	if err != nil {
		log.Fatalf("Cannot connect to Cloudinary: %v", err)
		return nil, err
	}
	return cld, nil
}
