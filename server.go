package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/context"

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
	mux.HandleFuncC(pat.Get("/bundle.js"), jsBundle)
	mux.HandleFuncC(pat.Get("/"), indexPage)

	loggedMux := handlers.LoggingHandler(os.Stdout, mux)

	http.ListenAndServe(listenStr, loggedMux)
}

func indexPage(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.IndexHtml)
}

func jsBundle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	f, err := os.Open("./assets/bundle.js")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	io.Copy(w, f)
}
