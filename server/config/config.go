package config

import (
	"os"
	"server/types"
	"strconv"
)

func LoadConfig() *types.Config {
	portValue := os.Getenv("PORT")

	portInt, _ := strconv.Atoi(portValue)

	return &types.Config{
		SecretKey:      os.Getenv("SECRET_KEY"),
		MaxConnections: 10,
		PORT:           portInt,
	}
}
