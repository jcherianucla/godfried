package config

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func StartRouter() (n *negroni.Negroni) {
	r := mux.NexRouter()
	n := negroni.Classic()
	n.UseHandler(r)

	return
}
