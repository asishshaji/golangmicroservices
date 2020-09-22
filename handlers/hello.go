package handlers

import (
	"fmt"
	"log"
)

type Hello struct {
	l *log.Logger
}

func (h *Hello) ServeHTTP() {
	h.l.Println("Hello")
}

func SayHello() {
	fmt.Println("Hety")
}
