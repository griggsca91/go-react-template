package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./templates/index.html")

	t.Execute(w, nil)
}

func main() {
	router := httprouter.New()
	router.GET("/", index)

	router.NotFound = http.FileServer(http.Dir("public"))

	err := http.ListenAndServe(":80", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	log.Print("Running on port 80")
}
