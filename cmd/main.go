package main

import (
	"fmt"
	"net/http"

	gen "github.com/Adambombtastic/grandgame/gen/game"
	genserver "github.com/Adambombtastic/grandgame/gen/http/game/server"
	"github.com/Adambombtastic/grandgame/service"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	s := service.New()

	endpoints := gen.NewEndpoints(s)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder

	svr := genserver.New(endpoints, mux, dec, enc, nil, nil)

	genserver.Mount(mux, svr)

	httpsvr := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Printf("Server listening on %s\n", httpsvr.Addr)
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}

}
