package userapi

import (
	"encoding/json"
	"net/http"
	"task-management/api/middleware"
	"task-management/apperror"
	"task-management/consts"
	"task-management/model"
	"task-management/util"
	"time"
)

// @Tags        User
// @Summary     user registration
// @Description user registration
// @Param       request body model.RegisterRequest true "register request"
// @Produce     json
// @Router      /register [post]
// @Success     201 {object} model.RegisterResponse "user register successfully"
func (a *api) RegisterRequest(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	var payload model.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return apperror.ErrBadRequest.Customize(err.Error()).LogWithLocation()
	}

	_, err = a.App.UserService.UserRegister(payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(util.SetResponse(nil, 1, "user register successfully."))
	return nil
}

// @Tags        User
// @Summary     user login
// @Description user login
// @Param       request body model.LoginRequest true "login request"
// @Produce     json
// @Router      /login [post]
// @Success     200 {object} model.LoginResponse "user login successfully"
func (a *api) Login(ctx *middleware.Context, w http.ResponseWriter, r *http.Request) error {
	var payload model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return apperror.ErrBadRequest.Customize(err.Error()).LogWithLocation()
	}

	userID, err := a.App.UserService.VerifyUser(payload)
	if err != nil {
		return err
	}

	token, err := a.App.JwtService.CreateJWTToken(userID, time.Hour*24, consts.JwtKey)
	if err != nil {
		return err
	}

	res := make(map[string]interface{})
	res["token"] = token.Value

	json.NewEncoder(w).Encode(util.SetResponse(res, 1, "user login successfully"))
	return nil
}
