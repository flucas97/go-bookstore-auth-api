package service

import (
	"github.com/flucas97/bookstore/auth-api/src/domain/access_token"
	"github.com/flucas97/bookstore/auth-api/src/utils/errors_utils"
)

// related with controller
type ServiceInterface interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

// related with Repository (database)
type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

type service struct {
	dbRepo Repository
}

// wich repository the service will use
func NewService(repo Repository) ServiceInterface {
	return &service{
		dbRepo: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors_utils.RestErr) {
	AccessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return AccessToken, nil
}