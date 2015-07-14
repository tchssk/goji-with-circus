package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	goji.Get("/", hello)

	fd := flag.Uint("fd", 0, "File descriptor to listen and serve.")
	flag.Parse()

	if *fd != 0 {
		listener, err := net.FileListener(os.NewFile(uintptr(*fd), ""))
		if err != nil {
			panic(err)
		}
		goji.ServeListener(listener)
	} else {
		goji.Serve()
	}
}
