package active

import (
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"github.com/FernandoCagale/c4-active/pkg/infra/logger"
)

//Service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//FindAll canais
func (s *Service) FindAll(query *entity.Query) (values []*entity.Active, count int, err error) {
	values, count, err = s.repo.FindAll(query)
	if err != nil {
		logger.WithFields(logger.Fields{"Message": err.Error()}).Error("[UseCase(FindAll)]FindAll")
	}
	return values, count, err
}
