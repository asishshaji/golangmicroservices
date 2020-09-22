package main

import (
	"net/http"

	"github.com/asishshaji/restapi/handlers"
)

func main() {
	handlers.SayHello()
	http.ListenAndServe(":9090", nil)
}
