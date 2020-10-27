package handler

import (
	"strconv"
	"time"

	"todo/model"
	"todo/storage"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("ThisIsASecretKey")

//UserSignUp register new user
func UserSignUp(c *gin.Context, storage *storage.UserDB) (*model.UserSignUpResult, error) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		return nil, err
	}
	password := []byte(user.Password)
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	createdUser, err := storage.CreateUser(user)
	if err != nil {
		return nil, err
	}
	// return createdUser, nil
	userSignUpResult := &model.UserSignUpResult{
		Email: createdUser.Email,
	}
	return userSignUpResult, nil
}

//UserLogin Login user
func UserLogin(c *gin.Context, storage *storage.UserDB) (*model.UserLoginResult, error) {
	userLogin := model.UserLogin{}
	if err := c.ShouldBind(&userLogin); err != nil {
		return nil, err
	}
	user, err := storage.LoginUser(userLogin.Email)
	if err != nil {
		return nil, err
	}
	// JWT
	password := []byte(userLogin.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		return nil, err
	}
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		Id:        strconv.Itoa(int(user.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecretKey)
	userLoginResult := &model.UserLoginResult{
		UserID: user.ID,
		Token:  tokenString,
	}
	c.SetCookie("Token", tokenString, 3600*24*365, "/", "", false, true)
	return userLoginResult, err
}
