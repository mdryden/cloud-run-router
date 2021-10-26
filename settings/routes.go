package settings

import "log"

type CloudRun struct {
	Service string `json:"service"`
	Secure  bool   `json:"secure"`
	// TODO: project id
	// TODO: region
}

type CloudFunction struct {
	Name   string `json:"name"`
	Secure bool   `json:"secure"`

	// TODO: project id
	// TODO: region
}

type Route struct {
	Prefix   string
	Run      *CloudRun      `json:"run,omitempty"`
	Function *CloudFunction `json:"function,omitempty"`
}

func writeInvalidRouteWarning(route Route, reason string) {
	log.Printf("WARNING: Route with prefix '%s' %s.  Route will be ignored.", route.Prefix, reason)
}

func (route Route) validateRoute() bool {
	if route.Function != nil && route.Run != nil {
		writeInvalidRouteWarning(route, "has more than one handler configured")
		return false
	}

	if route.Function == nil && route.Run == nil {
		writeInvalidRouteWarning(route, "has no handler configured")
		return false
	}

	if route.Run != nil && route.Run.Service == "" {
		writeInvalidRouteWarning(route, "has no cloud run service configured")
		return false
	}

	if route.Function != nil && route.Function.Name == "" {
		writeInvalidRouteWarning(route, "has no function name configured")
		return false
	}

	if route.Prefix == "" {
		log.Print("WARNING: Route with missing prefix encountered.  Route will be ignored.")
		return false
	}
	// TODO: all kinds of other stuff to clean and validate the prefix, I'm sure.
	// Disallow routes starting with _
	// Remove any trailing slashes
	// Regex for valid url characters?

	return true
}
