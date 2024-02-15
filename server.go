package main

import (
	"fmt"
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

func FastHttpServer(handler fasthttp.RequestHandler) {
	server := &fasthttp.Server{
		Handler: handler,
	}
	listen, err := net.Listen("tcp", APP_URL)
	if err != nil {
		log.Fatalf("Error[net-listen]: %v", err)
	}

	fmt.Printf("Server: %s", APP_URL)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Error[fasthttp-server]: %v", err)
	}
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Error[shutting-down-server]: while shutting down the server: %v", err)
	}
}
