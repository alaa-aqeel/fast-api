package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/valyala/fasthttp"
)

const APP_HOST = "127.0.0.1"
const APP_PORT = 8080

func HandlerReqeust(ctx *fasthttp.RequestCtx) {

}

func main() {
	server := &fasthttp.Server{
		Handler: HandlerReqeust,
	}
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatalf("Error[net-listen]: %v", err)
	}

	fmt.Printf("Server: http://%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Error[fasthttp-server]: %v", err)
	}
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Error[shutting-down-server]: while shutting down the server: %v", err)
	}

}
