package repository

import (
	"errors"

	"github.com/Purvig648/graphql-demo/models"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}
type UserRepo interface {
	CreateUser(userDetails models.User) (models.User, error)
	CreateCompany(companyDetails models.Company) (models.Company, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
