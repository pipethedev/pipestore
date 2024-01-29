package types

import "pipebase/server/enums"

type AuthRequestStruct struct {
	Auth struct {
		Username string `json:"username"`
		APIKey   string `json:"apiKey"`
	} `json:"auth"`
}

type RecordRequestStruct struct {
	Type   enums.Operation
	Record struct{}
}
