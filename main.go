package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
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
	for _, s := range signals.ForwardSignals {
		fmt.Printf("%+v\n", s)
	}

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
