package app

import (
	"task-management/app/jwtauth"
	"task-management/app/task"
	"task-management/app/user"
	"task-management/model"
	"task-management/mongodatabase"
)

type App struct {
	Repos       *model.Repos
	TaskService task.Service
	UserService user.Service
	JwtService  jwtauth.Service
}

func New() (*App, error) {

	mongoDBConf, err := mongodatabase.InitConfig()
	if err != nil {
		return nil, err
	}

	mongoDBClient, err := mongoDBConf.NewConnection()
	if err != nil {
		return nil, err
	}

	repos := &model.Repos{
		MongoDBClient: mongoDBClient,
		MongoDB:       mongoDBConf,
	}

	appobject := App{
		Repos:       repos,
		TaskService: task.NewService(repos),
		UserService: user.NewService(repos),
		JwtService:  jwtauth.NewService(),
	}

	return &appobject, nil
}
