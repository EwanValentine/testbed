package config

import (
	"context"
	"log"

	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	Addr string `env:"ADDR,default=:9000"`
}

func Load() {
	var c Config
	ctx := context.Background()
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
}
