package main

import (
	"net/http"

	"github.com/asishshaji/restapi/handlers"
)

// sadasd
func main() {
	handlers.SayHello()
	http.ListenAndServe(":9090", nil)
}
