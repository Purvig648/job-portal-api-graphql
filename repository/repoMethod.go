package repository

import (
	"errors"

	"github.com/Purvig648/graphql-demo/models"
)

func (r *Repo) CreateUser(userDetails models.User) (models.User, error) {
	result := r.DB.Create(&userDetails)
	if result.Error != nil {
		return models.User{}, errors.New("could not create the records")
	}
	return userDetails, nil
}
func (r *Repo) CreateCompany(companyDetails models.Company) (models.Company, error) {
	result := r.DB.Create(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not create the records")
	}
	return companyDetails, nil

}
