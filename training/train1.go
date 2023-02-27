package main

import (
	"context"
	"fmt"
	"net/http"
	"training/global"
	"training/handlers"
	"training/models"
	"training/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	dbSession := services.GetNewDbSession()
	dbSession.AddUser(&models.User{
		ID:   1,
		Name: "Christy",
		From: "ijk",
	})
	global.GlobalCtx = context.WithValue(global.GlobalCtx, "db", dbSession)

	r := chi.NewRouter() //new router
	r.Use(middleware.Logger)
	r.Use(handlers.Middleware)
	fmt.Println("service started")
	// fmt.Println("service", dbSession)
	r.Get("/user", handlers.GetUser)
	r.Post("/add", handlers.AddUser)
	r.Put("/update", handlers.UpdateUser)
	r.Delete("/delete", handlers.DeleteUser)

	http.ListenAndServe(":3000", r) //router to start the server
}
