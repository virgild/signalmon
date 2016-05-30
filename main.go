package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		for {
			select {
			case t := <-time.After(time.Second):
				if t.Second() == 0 {
					makeReadings()
				}
			}
		}
	}()

	go StartServer(3000)

	for s := range ch {
		switch s {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGTERM:
			fallthrough
		case syscall.SIGKILL:
			fallthrough
		case syscall.SIGINT:
			return
		}
	}
}

func makeReadings() {
	url := "http://192.168.100.1/Diagnostics.asp"
	body, err := fetchPage(url)
	if err != nil {
		panic(err)
	}

	signals, err := ParseDiagnosticsPage(body)
	if err != nil {
		fmt.Printf("Could not parse page result - %v", err)
		os.Exit(1)
	}
	// for _, s := range signals.ForwardSignals {
	// 	fmt.Printf("%+v\n", s)
	// }

	db, err := InitDB("readings.db")
	if err != nil {
		panic(err)
	}

	err = InsertData(db, signals)
	if err != nil {
		panic(err)
	}
}

func fetchPage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(contents), nil
	}
}
