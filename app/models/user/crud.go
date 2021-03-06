package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/types"
)

func (u *User) Create() (err error) {
	if err = model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(uid string) (User, error) {
	var user User
	id := types.StringToInt(uid)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
