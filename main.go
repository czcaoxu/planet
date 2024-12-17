package main

import (
	"fmt"
	"net/http"
	"plat/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}

	server.ListenAndServe()
	fmt.Println("server start")
}
