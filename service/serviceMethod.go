package service

import (
	"strconv"

	"github.com/Purvig648/graphql-demo/graph/model"
	"github.com/Purvig648/graphql-demo/models"
	"github.com/Purvig648/graphql-demo/pkg"
)

func (s *Service) CreateCompany(companyDetails model.NewCompnay) (*model.Company, error) {
	cd := models.Company{
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}
	cd, err := s.userRepo.CreateCompany(cd)
	if err != nil {
		return nil, err
	}
	cid := strconv.FormatUint(uint64(cd.ID), 10)
	return &model.Company{
		ID:        cid,
		Name:      cd.Name,
		Location:  cd.Location,
		CreatedAt: cd.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: cd.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *Service) UserSignup(userData model.NewUser) (*model.User, error) {
	hashedPassword, err := pkg.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userDetails := models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		HashPassword: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Name:      userDetails.Name,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
