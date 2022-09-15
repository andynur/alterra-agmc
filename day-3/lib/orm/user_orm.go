package orm

import (
	"errors"
	"rest-echo/config"
	"rest-echo/lib/dto"
	"rest-echo/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserOrm struct{}

func (o *UserOrm) GetAllUser() (users []models.User, err error) {
	if err := config.DB.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (o *UserOrm) FindUserById(id int) (user models.User, err error) {
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (o *UserOrm) FindUserByEmailAndPassword(email string, password string) (user *models.User, err error) {
	if err := config.DB.First(&user, "email = ?", email).Error; err != nil {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("incorrect password")
	}

	return user, nil
}

func (o *UserOrm) CreateUser(req dto.StoreUserRequest) (user models.User, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
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
		return user, err
	}

	return storeUser, nil
}

func (o *UserOrm) UpdateUserById(id int, req dto.UpdateUserRequest) (user models.User, err error) {
	user, err = o.FindUserById(id)
	if err != nil {
		return user, err
	}

	password := user.Password
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return user, err
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
	DB := config.DB
	if err := DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(&updateUser).Error; err != nil {
		return user, err
	}

	return updateUser, nil
}

func (o *UserOrm) DeleteUserById(id int) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
