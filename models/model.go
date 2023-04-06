package models

type ElasticErrorModel struct {
	Error  ElasticError `json:"error"`
	Status int          `json:"status"`
}

type ElasticError struct {
	Reason string `json:"reason"`
	Index  string `json:"index"`
	Type   string `json:"type"`
}
