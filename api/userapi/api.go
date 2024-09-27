package userapi

import (
	"task-management/app"
)

type api struct {
	App *app.App
}

// New creates a new api
func New(app *app.App) *api {
	return &api{
		App: app,
	}
}
