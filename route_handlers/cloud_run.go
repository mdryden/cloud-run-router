package route_handlers

import "net/http"

type CloudRun struct {
	Service string `json:"service"`
	Secure  bool   `json:"secure"`
	// TODO: project id
	// TODO: region
}

func (CloudRun) HandleRequest(w http.ResponseWriter, r http.Request) {

}
