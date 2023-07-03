package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"grcp"`
	Port    int32  `default:"8000"`
	DB      struct {
		Use   string `default:"postgres"`
		Postgres []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"xxxxxxxx"`
			Database string `default:"Majordome"`
		}
	}
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
