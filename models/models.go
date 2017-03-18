package models

type InRequest struct {
	Source Source `json:"source"`
}

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutParams struct {
	Directory string `json:"directory"`
}

type Source struct {
	URI           string `json:"uri"`
	CaCert        string `json:"ca_cert"`
	Authorization string `json:"authorization"`
}

type InResponse struct{}

type CheckRequest struct{}

type CheckResponse []interface{}
