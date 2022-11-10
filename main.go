package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

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

	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	server.ListenAndServe()
}
