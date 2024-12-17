package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"plat/framework"
	"plat/framework/middleware"
	"syscall"
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
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	fmt.Println("server shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
}
