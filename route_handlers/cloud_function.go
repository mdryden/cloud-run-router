package route_handlers

import "net/http"

type CloudFunction struct {
	Name   string `json:"name"`
	Secure bool   `json:"secure"`

	// TODO: project id
	// TODO: region
}

func (CloudFunction) HandleRequest(w http.ResponseWriter, r http.Request) {

}
