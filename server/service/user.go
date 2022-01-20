package service

import (
	"context"
	"errors"

	"github.com/arganaphangquestian/gerpisi/server/data"
	"github.com/arganaphangquestian/gerpisi/server/utils"
	"github.com/segmentio/ksuid"
)

type UserService struct{}

type User struct {
	Id       *string
	Name     string
	Email    string
	Password string
}

var USERS = []User{}

func (s *UserService) Login(ctx context.Context, request *data.LoginRequest) (*data.LoginResponse, error) {
	var user *data.User
	for _, v := range USERS {
		if v.Email == request.Email && utils.CompareHash(v.Password, request.Password) {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	token, err := utils.CreateToken(*user.Id)
	if err != nil {
		return nil, errors.New("can't create token")
	}
	return &data.LoginResponse{
		AccessToken: token.AccessToken,
	}, nil
}

func (s *UserService) AddUser(ctx context.Context, request *data.AddUserRequest) (*data.AddUserResponse, error) { // Protected Route/endpoint
	// Check JWT
	userId, err := utils.Verify(request.GetAccessToken())
	if err != nil {
		return nil, errors.New("token invalid")
	}
	var user *data.User
	for _, v := range USERS {
		if v.Id == userId {
			user = &data.User{
				Id:       v.Id,
				Name:     v.Name,
				Email:    v.Email,
				Password: v.Password,
			}
			break
		}
	}
	if user == nil {
		return nil, errors.New("un-authorize")
	}
	id := ksuid.New().String()
	pass, _ := utils.CreateHash(request.User.Password)
	newUser := User{
		Id:       &id,
		Name:     request.User.Name,
		Email:    request.User.Email,
		Password: *pass,
	}
	USERS = append(USERS, newUser)
	return &data.AddUserResponse{
		Message: "WokAy",
		User: &data.User{
			Id:       newUser.Id,
			Name:     newUser.Name,
			Email:    newUser.Email,
			Password: newUser.Password,
		},
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, request *data.GetUsersRequest) (*data.GetUsersResponse, error) {
	var users []*data.User
	for _, v := range USERS {
		users = append(users, &data.User{
			Id:       v.Id,
			Name:     v.Name,
			Email:    v.Email,
			Password: v.Password,
		})
	}
	return &data.GetUsersResponse{
		Message: "Get all data",
		Users:   users,
	}, nil
}
