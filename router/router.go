package router

import (
	"backend/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/login", handler.Login)

	fmt.Println("Listening and serving on http://localhost:4000")
	err := http.ListenAndServe(":4000", r)
	if err != nil {
		fmt.Println(err)
	}
}
