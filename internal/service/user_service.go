package service

import (
	"errors"
	"strings"

	"github.com/nunusavi/task-manager/internal/model"
	"github.com/nunusavi/task-manager/internal/respository"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string)(*model.User, error){
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	existing, _ := respository.GetUserByEmail(email)
	if existing != nil{
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return nil, errors.New("failed to has password")
	}

	user := &model.User{
		Email: email,
		Password: string(hashed),
	}

	err = respository.CreateUser(user)
	if err != nil{
		return nil, errors.New("could not save user")
	}

	return user, nil
}