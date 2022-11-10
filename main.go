package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Get")
	})

	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Products Id : %s", p.ByName("id"))
	})

	router.GET(
		"/products/:id/items/:itemsId",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Fprintf(w, "Product Id : %s\nItems Id : %s", p.ByName("id"), p.ByName("itemsId"))
		},
	)

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Image Path : %s", p.ByName("image"))
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}
