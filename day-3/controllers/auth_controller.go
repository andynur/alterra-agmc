package controllers

import (
	"net/http"
	"rest-echo/lib/dto"
	"rest-echo/lib/orm"
	"rest-echo/lib/utils"

	"github.com/labstack/echo/v4"
)

func AuthLoginController(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	userOrm := new(orm.UserOrm)
	user, err := userOrm.FindUserByEmailAndPassword(req.Email, req.Password)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	jwtUtils := new(utils.JwtUtils)
	tokenData, err := jwtUtils.GenerateToken(user.ID)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, tokenData)
}
