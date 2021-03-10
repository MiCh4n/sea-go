package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Response string
	Color    bool
}

func page(data string, pirates bool) {
	_Request := Request{
		Response: data,
		Color:    pirates}
	LISTEN_ADDRESS := os.Getenv("LISTEN_ADDRESS")
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, _Request)
	})

	log.Fatal(http.ListenAndServe(LISTEN_ADDRESS, nil))
}
func main() {
	UPSTREAM_ADDRESS := os.Getenv("UPSTREAM_ADDRESS")
	time.Sleep(time.Second * 10)
	client := http.Client{}
	req, err := http.NewRequest("GET", UPSTREAM_ADDRESS, nil)

	response, err := client.Do(req)
	if err != nil {
		c := true
		page("There are no pirates on the sea! Pods are safe!", c)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := false
	page(string(responseData), c)

}
