package service

import (
	"errors"

	"github.com/Purvig648/graphql-demo/graph/model"
	"github.com/Purvig648/graphql-demo/repository"
)

type UserService interface {
	UserSignup(userData model.NewUser) (*model.User, error)
	CreateCompany(companyDetails model.NewCompnay) (*model.Company, error)
}
type Service struct {
	userRepo repository.UserRepo
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be nil")
	}
	return &Service{
		userRepo: userRepo,
	}, nil

}
