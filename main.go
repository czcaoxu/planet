package main

import (
	"fmt"
	"net/http"
	"plat/framework"
	"plat/framework/middleware"
)

func main() {
	//log.SetFlags(log.Llongfile)
	core := framework.NewCore()
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)

	server := &http.Server{
		Addr:    ":8080",
		Handler: core,
	}
	fmt.Println("server start")
	server.ListenAndServe()
}
