package services

import (
	"github.com/Cahllagerfeld/go-htmx-first-steps/internal/domain"
	"github.com/markbates/goth"
)

var users = []*domain.User{}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) FindOrCreateUser(user goth.User) *domain.User {
	for _, v := range users {
		if v.Email == user.Email {
			return v
		}
	}
	u := &domain.User{
		ID:       len(users) + 1,
		Email:    user.Email,
		Username: user.NickName,
		Avatar:   user.AvatarURL,
	}
	users = append(users, u)
	return u
}
