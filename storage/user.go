package storage

import (
	"../model"
	"github.com/jinzhu/gorm"
)

//UserDBLayer is the interface for user database
type UserDBLayer interface {
	CreateUser(model.User) (*model.User, error)
	LoginUser(string) (*model.User, error)
}

//UserDB is the layer for database
type UserDB struct {
	DB *gorm.DB
}

//CreateUser create a new user
func (u *UserDB) CreateUser(user model.User) (*model.User, error) {
	err := u.DB.Create(&user).Error
	return &user, err
}

//LoginUser use email
func (u *UserDB) LoginUser(email string) (*model.User, error) {
	user := &model.User{}
	err := u.DB.
		Where("email = ?", email).
		First(user).Error
	return user, err
}
