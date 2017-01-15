package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/virgild/signalmon/templates"
	"goji.io"
	"goji.io/pat"
)

// port := flag.Int("p", 8000, "Listen port number")
// flag.Parse()

// fmt.Fprint(os.Stdout, fmt.Sprintf("Phleeeeeep running at port %v:\n", *port))

// listenStr := fmt.Sprintf(":%v", *port)

// mux := goji.NewMux()
// mux.HandleFuncC(pat.Get("/t"), t)

// loggedMux := handlers.LoggingHandler(os.Stdout, mux)

// go runFetcher()
// startFetch()

// http.ListenAndServe(listenStr, loggedMux)

func StartServer(port int) {
	listenStr := fmt.Sprintf(":%d", port)
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/bundle.js"), jsBundle)
	mux.HandleFunc(pat.Get("/"), indexPage)

	loggedMux := handlers.LoggingHandler(os.Stdout, mux)

	http.ListenAndServe(listenStr, loggedMux)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.IndexHtml)
}

func jsBundle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/bundle.js")
}
