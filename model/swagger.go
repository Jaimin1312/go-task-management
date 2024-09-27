package model

type ServerError500 struct {
	Data    string `json:"data" example:"null"`
	Status  int    `json:"status" example:"0"`
	Message string `json:"message" example:"internal server error"`
}

type ServerError401 struct {
	Data    string `json:"data" example:"null"`
	Status  int    `json:"status" example:"0"`
	Message string `json:"message" example:"Unauthorized access"`
}

type TaskCreateResponse struct {
	Message string `json:"message" example:"Task created successfully."`
	Status  int    `json:"status" example:"1"`
	Data    string `json:"data" example:"60d5ec49c6d8c06e1f20c5a8"`
}

type TaskReadResponse struct {
	Message string `json:"message" example:"Task retrieved successfully."`
	Status  int    `json:"status" example:"1"`
	Data    Task   `json:"data"`
}

type TaskListResponse struct {
	Message string `json:"message" example:"Task list retrieved successfully."`
	Status  int    `json:"status" example:"1"`
	Data    []Task `json:"data"`
}

type TaskUpdateResponse struct {
	Data    string `json:"data" example:"null"`
	Status  int    `json:"status" example:"1"`
	Message string `json:"message" example:"Task updated successfully"`
}

type TaskDeleteResponse struct {
	Data    string `json:"data" example:"null"`
	Status  int    `json:"status" example:"1"`
	Message string `json:"message" example:"Task deleted successfully"`
}

type MarkTasksAsDoneResponse struct {
	Data    []string `json:"data" example:"Task 60d5ec49c6d8c06e1f20c5a8 marked as done."`
	Status  int      `json:"status" example:"1"`
	Message string   `json:"message" example:"Tasks processed."`
}

type LoginResponse struct {
	Data struct {
		Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOiI2NmY1ODY1NjhiYTAxYjBkOGQ4MzFjMjUiLCJleHAiOjUxNjc2MzY4NzB9.s4U-8Hl6s3hTr0n0Zb9FbYLEGCwO4k5pL5trZxd6AeI"`
	} `json:"data"`
	Message string `json:"message" example:"user login successfully"`
	Status  int    `json:"status" example:"1"`
}

type RegisterResponse struct {
	Data    string `json:"data" example:"null"`
	Status  int    `json:"status" example:"1"`
	Message string `json:"message" example:"user register successfully"`
}
