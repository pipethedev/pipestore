package types

import "pipebase/server/enums"

type AuthRequestStruct struct {
	Auth struct {
		Username string `json:"username"`
		APIKey   string `json:"apiKey"`
	} `json:"auth"`
}

type RecordRequestStruct struct {
	Data struct {
		Type      enums.Operation `json:"type"`
		TableName string          `json:"tableName"`
		Record    interface{}     `json:"record"`
	} `json:"data"`
}
