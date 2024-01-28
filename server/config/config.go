package config

import (
	"os"
	"pipebase/server/types"
)

func LoadConfig() *types.Config {
	return &types.Config{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
