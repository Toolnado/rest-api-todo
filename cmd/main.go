package main

import (
	"log"

	todo "github.com/toolnado/rest-todo-app"
	"github.com/toolnado/rest-todo-app/pkg/handler"
)

func main() {
	svr := new(todo.Server)
	handlers := new(handler.Handler)
	if err := svr.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
