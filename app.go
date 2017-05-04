package main

import (
	"fmt"
	"github.com/jcherianucla/godfried/config"
	"net/http"
)

func main() {
	r := config.StartRouter()
	port := config.GetPort()
	http.ListenAndServe(port, n)
}
