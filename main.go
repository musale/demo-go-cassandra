package main

import (
	"net/http"

	"github.com/musale/demo-go-cassandra/lib"
)

func main() {
	http.HandleFunc("/messages/send", lib.SendMessage)
	http.ListenAndServe(":5050", nil)
}
