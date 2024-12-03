package todo_server

import (
	"log"
	"net/http"

	todoshandler "examles_11_cohort/todo_server/internal/handler"
	"examles_11_cohort/todo_server/internal/repository"
	"examles_11_cohort/todo_server/pkg"
)

func RunServer() {
	repo := repository.NewInMemoryRepo()
	handler := todoshandler.NewTodosHandler(repo)

	http.Handle("/todos", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.GetTodos))))
	http.Handle("/add", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.AddTodo))))
	http.Handle("/update", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.UpdateTodo))))
	http.Handle("/delete", pkg.CorsMiddleware(pkg.LoggingMiddleware(http.HandlerFunc(handler.DeleteTodo))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("can not start the server: %+v", err)
	}
}
