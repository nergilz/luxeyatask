package config

import "time"

type Config struct {
	Address string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}
