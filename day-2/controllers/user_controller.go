package controllers

import (
	"errors"
	"net/http"
	"rest-echo/config"
	"rest-echo/lib"
	"rest-echo/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersController(c echo.Context) error {
	DB := config.DB
	var users []models.User
	if err := DB.Find(&users).Error; err != nil {
		return ErrorResponse(c, err, http.StatusInternalServerError)
	}

	return SuccessResponse(c, users)
}

func ShowUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	if userId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	DB := config.DB
	var user models.User
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, user)
}

func StoreUserController(c echo.Context) error {
	var req lib.StoreUserRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	storeUser := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    string(hashedPassword),
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}

	DB := config.DB
	if err = DB.Create(&storeUser).Error; err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, storeUser)
}

func UpdateUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	if userId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	var req lib.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	if err := c.Validate(&req); err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	DB := config.DB
	var user models.User
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	password := user.Password
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return ErrorResponse(c, err, http.StatusBadRequest)
		}

		password = string(hashedPassword)
	}

	updateUser := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    password,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}

	updateUser.ID = user.ID
	updateUser.UpdatedAt = time.Now()
	if err := DB.Model(&models.User{}).Where("id = ?", userId).Updates(&updateUser).Error; err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, updateUser)
}

func DeleteUserController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	if userId == 0 {
		return ErrorResponse(c, errors.New("id path variables is required"), http.StatusBadRequest)
	}

	DB := config.DB
	if err := DB.Delete(&models.User{}, userId).Error; err != nil {
		return ErrorResponse(c, err, http.StatusBadRequest)
	}

	return SuccessResponse(c, nil)
}
