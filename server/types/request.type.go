package types

type RequestStruct struct {
	Request struct{} `json:"request"`
	Auth    struct {
		Username string `json:"username"`
		APIKey   string `json:"apiKey"`
	} `json:"auth"`
}
