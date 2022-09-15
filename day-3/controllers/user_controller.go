package controllers

import (
	"errors"
	"net/http"
	"rest-echo/lib/dto"
	"rest-echo/lib/orm"
	"rest-echo/lib/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	userOrm := new(orm.UserOrm)
	users, err := userOrm.GetAllUser()
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, users)
}

func ShowUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	if userId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	userOrm := new(orm.UserOrm)
	user, err := userOrm.FindUserById(userId)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, user)
}

func StoreUserController(c echo.Context) error {
	var req dto.StoreUserRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	userOrm := new(orm.UserOrm)
	user, err := userOrm.CreateUser(req)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, user)
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	// compare id with auth user id
	jwtUtils := new(utils.JwtUtils)
	userId := jwtUtils.ExtractTokenUserId(c)
	if userId != id {
		return ErrorResponse(c, errors.New("incorrect user id"), http.StatusForbidden)
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	userOrm := new(orm.UserOrm)
	user, err := userOrm.UpdateUserById(userId, req)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, user)
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	// compare id with auth user id
	jwtUtils := new(utils.JwtUtils)
	userId := jwtUtils.ExtractTokenUserId(c)
	if userId != id {
		return ErrorResponse(c, errors.New("incorrect user id"), http.StatusForbidden)
	}

	userOrm := new(orm.UserOrm)
	if err := userOrm.DeleteUserById(userId); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	// TODO: check blacklist token on redis

	return SuccessResponse(c, nil)
}
