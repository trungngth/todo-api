package model

import "github.com/jinzhu/gorm"

//User with email and password
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	//Notes    []Note `json:"notes" gorm:"foreignkey:user_id;type:varchar"`
}

//UserLogin stores login information
type UserLogin struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

//UserSignUpResult return when signup successfully
type UserSignUpResult struct {
	Email string
}

//UserLoginResult returns userid and jwt
type UserLoginResult struct {
	UserID uint
	Token  string
}
