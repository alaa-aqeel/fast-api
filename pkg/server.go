package pkg

import (
	"fmt"
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

func FastHttpServer(handler fasthttp.RequestHandler, appUrl string) {
	server := &fasthttp.Server{
		Handler: handler,
	}
	listen, err := net.Listen("tcp", appUrl)
	if err != nil {
		log.Fatalf("Error[net-listen]: %v", err)
	}

	fmt.Printf("Server: %s", appUrl)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Error[fasthttp-server]: %v", err)
	}
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Error[shutting-down-server]: while shutting down the server: %v", err)
	}
}
