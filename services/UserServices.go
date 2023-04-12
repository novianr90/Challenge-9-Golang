package services

import (
	"challenge-9/models"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) GetAllUser() ([]models.User, error) {
	var users []models.User
	if err := us.DB.Find(&users).Error; err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (us *UserService) CreateUser(user models.User) (models.User, error) {
	if err := us.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := us.DB.Where("email = ?", email).First(&user).Error

	if err != nil || user.Email == "" || user.Password == "" {
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) UpdateUserByEmail(email string, user models.User) error {
	result := us.DB.Model(models.User{}).Where("email = ?", email).Updates(&user)

	if result.RowsAffected == 0 {
		return errors.New("there's no data to update")
	}
	return nil
}

func (us *UserService) DeleteUserByEmail(email string) error {
	var user models.User

	if err := us.DB.Where("email = ?", email).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
