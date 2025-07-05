package main

import (
	handlers "_/D_/MicroServ"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	http.ListenAndServe(":9090", nil)
}
