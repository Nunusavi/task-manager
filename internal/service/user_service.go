package service

import (
	"errors"
	"strings"

	"github.com/nunusavi/task-manager/internal/model"
	"github.com/nunusavi/task-manager/internal/repository"
	"github.com/nunusavi/task-manager/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string) (*model.User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	existing, _ := repository.GetUserByEmail(email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to has password")
	}

	user := &model.User{
		Email:    email,
		Password: string(hashed),
	}

	err = repository.CreateUser(user)
	if err != nil {
		return nil, errors.New("could not save user")
	}

	return user, nil
}

func LoginUser(email, password string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(int64(user.ID))
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
