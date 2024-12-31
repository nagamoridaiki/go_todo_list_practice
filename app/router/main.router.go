package router

import (
	"fmt"
	"net/http"

	"myapp/middleware"

	"github.com/gorilla/mux"
)

type MainRouter interface {
	setupRouting() *mux.Router
	StartWebServer() error
}

type mainRouter struct {
	appR AppRouter

	todoR TodoRouter
}

func NewMainRouter(appR AppRouter, todoR TodoRouter) MainRouter {
	return &mainRouter{appR, todoR}
}

const PORT = 4000

// routing definition
func (mainRouter *mainRouter) setupRouting() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// CORSミドルウェアを追加
	router.Use(middleware.Cors)

	// OPTIONSメソッドを許可するためのハンドラーを追加
	router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mainRouter.appR.SetAppRouting(router)
	mainRouter.todoR.SetTodoRouting(router)

	return router
}

// StartWebServer server startup
func (mainRouter *mainRouter) StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")

	return http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter.setupRouting())
}
