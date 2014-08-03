package main

import (
	"net/http"
	"log"
	."logic/handlers"
)

func main() {
	log.Println("main")
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/index/", Index.IndexAction)

	http.ListenAndServe(":8888", nil)
}
