package taskapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"task-management/api/middleware"
	"task-management/apperror"
	"task-management/model"
	"task-management/util"

	"github.com/gorilla/mux"
)

// @Tags        Task
// @Summary     Create a task
// @Description Create a new task for a user
// @Produce     json
// @Param       auth-token header   string                   true "token value"
// @Param       payload    body     model.TaskCreateRequest  true "Task creation request"
// @Success     201        {object} model.TaskCreateResponse "Task created successfully"
// @Failure     500        {object} model.ServerError500
// @Failure     401        {object} model.ServerError401
// @Router      /task [post]
func (a *api) TaskCreate(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	var payload model.TaskCreateRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	// Validate the task
	if err := payload.Validate(); err != nil {
		return err
	}

	payload.UserID = ctx.UserID

	id, err := a.App.TaskService.TaskCreate(payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(util.SetResponse(id, 1, "Task created successfully."))
	return nil
}

// @Tags        Task
// @Summary     Retrieve a task by ID
// @Description Get the details of a task by its ID
// @Produce     json
// @Param       id         path     string                 true "Task ID"
// @Param       auth-token header   string                 true "token value"
// @Success     200        {object} model.TaskReadResponse "Task created successfully"
// @Failure     500        {object} model.ServerError500
// @Failure     401        {object} model.ServerError401
// @Router      /task/{id} [get]
func (a *api) TaskRead(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	taskID := params["id"]

	task, err := a.App.TaskService.TaskRead(taskID, ctx.UserID)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(util.SetResponse(task, 1, "Task retrieved successfully."))
	return nil
}

// @Tags        Task
// @Summary     List all tasks
// @Param       auth-token header string true "token value"
// @Description Retrieve all tasks for a user with optional filters
// @Produce     json
// @Param       limit     query    int                    false "Limit of tasks to retrieve"
// @Param       offset    query    int                    false "Offset for pagination"
// @Param       status    query    string                 false "Filter by task status"
// @Param       sort      query    string                 false "Field to sort by"
// @Param       sortOrder query    int                    false "Order of sorting (1 for ascending, -1 for descending)"
// @Success     200       {object} model.TaskListResponse "Task created successfully"
// @Failure     500       {object} model.ServerError500
// @Failure     401       {object} model.ServerError401
// @Router      /task [get]
func (a *api) TaskList(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	statusFilter := r.URL.Query().Get("status")
	sortField := r.URL.Query().Get("sort")
	sortOrder, _ := strconv.Atoi(r.URL.Query().Get("sortOrder"))

	var status *string
	if statusFilter != "" {
		status = &statusFilter
	}

	var sort *string
	if sortField != "" {
		sort = &sortField
	}

	tasks, err := a.App.TaskService.TaskList(ctx.UserID, limit, offset, status, sort, &sortOrder)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(util.SetResponse(tasks, 1, "Task list retrieved successfully."))
	return nil
}

// @Tags        Task
// @Summary     Update a task by ID
// @Description Update the details of a task by its ID
// @Param       auth-token header string true "token value"
// @Produce     json
// @Param       id      path     string                     true "Task ID"
// @Param       payload body     model.TaskUpdateRequest    true "Task update request"
// @Success     200     {object} model.TaskUpdateResponse   "Task updated successfully"
// @Failure     400     {object} util.Response{data=string} "Bad request"
// @Failure     500     {object} model.ServerError500
// @Failure     401     {object} model.ServerError401
// @Router      /task/{id} [put]
func (a *api) TaskUpdate(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	taskID := params["id"]

	var payload model.TaskUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	// Validate the task
	if err := payload.Validate(); err != nil {
		return err
	}

	err = a.App.TaskService.TaskUpdate(taskID, payload, ctx.UserID)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(util.SetResponse(nil, 1, "Task updated successfully."))
	return nil
}

// @Tags        Task
// @Summary     Delete a task by ID
// @Description Remove a task from the system by its ID
// @Param       auth-token header string true "token value"
// @Produce     json
// @Param       id  path     string                   true "Task ID"
// @Success     200 {object} model.TaskDeleteResponse "Task deleted successfully"
// @Failure     500 {object} model.ServerError500
// @Failure     401 {object} model.ServerError401
// @Router      /task/{id} [delete]
func (a *api) TaskDelete(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	taskID := params["id"]

	err := a.App.TaskService.TaskDelete(taskID, ctx.UserID)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(util.SetResponse(nil, 1, "Task deleted successfully."))
	return nil
}

// @Tags        Task
// @Summary     Mark multiple tasks as done
// @Description Mark multiple tasks as done for a user
// @Param       auth-token header string true "token value"
// @Produce     json
// @Param       payload body     model.MarkDoneRequest         true "Task IDs to mark as done"
// @Success     200     {object} model.MarkTasksAsDoneResponse "Tasks marked as done successfully"
// @Failure     400     {object} util.Response{data=string}    "Bad request"
// @Failure     500     {object} model.ServerError500
// @Failure     401     {object} model.ServerError401
// @Router      /tasks/mark-done [put]
func (a *api) MarkTasksAsDone(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	var payload model.MarkDoneRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return apperror.ErrBadRequest.Customize(err.Error()).LogWithLocation()
	}

	// Validate the task IDs
	if len(payload.TaskIDs) == 0 {
		return apperror.ErrBadRequest.Customize("task_ids cannot be empty") // Custom error for empty IDs
	}

	// Channel to collect results
	results := make(chan string, len(payload.TaskIDs))

	// WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Define a function to mark tasks as done
	markTask := func(taskID string) {
		defer wg.Done()
		if err := a.App.TaskService.MarkTasksAsDone([]string{taskID}, ctx.UserID); err != nil {
			results <- "Error marking task " + taskID + ": " + err.Error()
		} else {
			results <- "Task " + taskID + " marked as done."
		}
	}

	// Start Goroutines for each task ID
	for _, taskID := range payload.TaskIDs {
		wg.Add(1)
		go markTask(taskID)
	}

	// Wait for all Goroutines to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	var responseMessages []string
	for msg := range results {
		responseMessages = append(responseMessages, msg)
	}

	// Send success response
	json.NewEncoder(w).Encode(util.SetResponse(responseMessages, 1, "Tasks processed."))
	return nil
}
