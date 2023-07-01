package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ProvideConfig() (Config, error) {
	v := viper.New()

	setDefaults(v)

	_ = v.ReadInConfig()

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return Config{}, fmt.Errorf("viper failed to unmarshal app config: %w", err)
	}

	return c, nil
}
