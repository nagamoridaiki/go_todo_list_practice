package router

import (
	"myapp/controllers"

	"github.com/gorilla/mux"
)

type TodoRouter interface {
	SetTodoRouting(router *mux.Router)
}

type todoRouter struct {
	tc controllers.TodoController
}

func NewTodoRouter(tc controllers.TodoController) TodoRouter {
	return &todoRouter{tc}
}

func (tr *todoRouter) SetTodoRouting(router *mux.Router) {
	router.HandleFunc("/api/v1/todo", tr.tc.FetchAllTodos).Methods("GET")
	router.HandleFunc("/api/v1/todo/{id}", tr.tc.FetchTodoById).Methods("GET")

	router.HandleFunc("/api/v1/todo", tr.tc.CreateTodo).Methods("POST")
	router.HandleFunc("/api/v1/todo/{id}", tr.tc.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/v1/todo/{id}", tr.tc.UpdateTodo).Methods("PUT")
}
