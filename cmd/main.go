package main

import (
	"fmt"
	"net/http"
	"os"

	gen "github.com/Adambombtastic/grandgame/gen/game"
	genserver "github.com/Adambombtastic/grandgame/gen/http/game/server"
	"github.com/Adambombtastic/grandgame/service"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	// load the port from the environment variable
	// if not set, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := service.New()

	endpoints := gen.NewEndpoints(s)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	svr := genserver.New(endpoints, mux, dec, enc, nil, nil)

	genserver.Mount(mux, svr)

	httpsvr := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	fmt.Printf("Server listening on %s\n", httpsvr.Addr)
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}

}
