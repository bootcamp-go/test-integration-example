package main

import (
	"app/cmd/server/handlers"
	"app/internal/products/storage"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// env
	addr := os.Getenv("SERVER_ADDR")

	// dependencies
	st := storage.NewStorageProductsMap(make(map[int]storage.ProductAttributes), 0)
	hd := handlers.NewHandlerProducts(st)

	// server
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - routes
	rt.Post("/products", hd.Save)

	// run
	if err := http.ListenAndServe(addr, rt); err != nil {
		fmt.Println(err)
		return
	}
}