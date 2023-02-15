package rest

import (
	"net/http"
)

// Rest server routes
type TRestServerEndpoints struct {
	Name     string
	Endpoint string
	Type     string
	Callback func(w http.ResponseWriter, r *http.Request)
}

// Rest server configuration
type TRestServerConf struct {
	Name      string
	Listen    string
	Endpoints []TRestServerEndpoints
}
