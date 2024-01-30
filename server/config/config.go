package config

import (
	"os"
	"server/types"
	"strconv"
)

func LoadConfig() *types.Config {
	portValue := os.Getenv("PORT")
	if portValue == "" {
		// port, _ := utils.AvailablePort()
		// portValue = strconv.Itoa(port)
	}

	portInt, _ := strconv.Atoi(portValue)

	//maxConnections, _ := strconv.Atoi(os.Getenv("MAX_CONNECTION"))

	return &types.Config{
		SecretKey:      os.Getenv("SECRET_KEY"),
		MaxConnections: 10,
		PORT:           portInt,
	}
}
