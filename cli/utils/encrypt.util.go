package utils

import (
	"pipebase/cli/types"
)

func Encrypt(credentials types.UserCredentials, key string) (types.UserCredentials, error) {

	return types.UserCredentials{
		Username: credentials.Username,
		APIKey:   credentials.APIKey,
	}, nil
}

func Decrypt(credentials types.UserCredentials, key string) (types.UserCredentials, error) {
	return types.UserCredentials{
		Username: credentials.Username,
		APIKey:   credentials.APIKey,
	}, nil
}
