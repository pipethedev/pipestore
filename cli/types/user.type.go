package types

type UserCredentials struct {
	Username    string
	Password    string
	APIKey      string
	StoragePath string
}

type InitConfig struct {
	Username string
	PassKey  string
}
