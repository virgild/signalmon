package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/context"
)

var leepits [30]string
var catstream chan string = make(chan string)

func main() {
	url := "http://192.168.100.1/Diagnostics.asp"
	body, err := fetchPage(url)
	if err != nil {
		panic(err)
	}
	stats, err := ParseDiagnosticsPage(body)
	if err != nil {
		panic(err)
	}
	json, err := json.Marshal(stats)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))

	// db, err := initdb("readings.db")
	// if err != nil {
	// 	panic(err)
	// }

	// err = insertData(db, stats)
	// if err != nil {
	// 	log.Println(err)
	// }
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

func t(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Good")
}
