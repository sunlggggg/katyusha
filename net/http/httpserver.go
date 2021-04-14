package http 

import (
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
)

func Server() {
	log.Println("Start http server.")
	router := httprouter.New()
    router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
