package logic

import (
	"gin_forum/dao/mysql"
	"gin_forum/models"
	"gin_forum/pkg/jwt"
	"gin_forum/pkg/snowflake"
)



func SignUp(p *models.ParamSignUp) (err error) {

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {

		return nil, err
	}
	
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
