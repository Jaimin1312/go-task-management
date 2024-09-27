package api

import (
	"net/http"
	"task-management/api/middleware"
	"task-management/api/taskapi"
	"task-management/api/userapi"
	"task-management/app"
	"task-management/consts"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type API struct {
	App *app.App
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}

	return api, nil
}

func (a *API) Init(r *mux.Router) {

	config := middleware.MiddlewareConfig{
		CookieName:     consts.CookieName,
		MaxContentSize: 1000,
		JwtService:     a.App.JwtService,
	}

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	usernewapi := userapi.New(a.App)
	r.HandleFunc("/register", config.MiddlewareHandler(usernewapi.RegisterRequest, false)).Methods(http.MethodPost)
	r.HandleFunc("/login", config.MiddlewareHandler(usernewapi.Login, false)).Methods(http.MethodPost)

	tasknewapi := taskapi.New(a.App)
	r.HandleFunc("/task", config.MiddlewareHandler(tasknewapi.TaskCreate, true)).Methods(http.MethodPost)               // Create task
	r.HandleFunc("/task", config.MiddlewareHandler(tasknewapi.TaskList, true)).Methods(http.MethodGet)                  // List tasks
	r.HandleFunc("/task/{id}", config.MiddlewareHandler(tasknewapi.TaskRead, true)).Methods(http.MethodGet)             // Read task by ID
	r.HandleFunc("/task/{id}", config.MiddlewareHandler(tasknewapi.TaskUpdate, true)).Methods(http.MethodPut)           // Update task by ID
	r.HandleFunc("/task/{id}", config.MiddlewareHandler(tasknewapi.TaskDelete, true)).Methods(http.MethodDelete)        // Delete task by ID
	r.HandleFunc("/task/mark-done", config.MiddlewareHandler(tasknewapi.MarkTasksAsDone, true)).Methods(http.MethodPut) // Mark multiple task done

}
