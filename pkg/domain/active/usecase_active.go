package active

import (
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"github.com/FernandoCagale/c4-active/pkg/infra/errors"
)

type ActiveUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *ActiveUseCase {
	return &ActiveUseCase{
		repo: repo,
	}
}

func (usecase *ActiveUseCase) Create(e *entity.Active) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	return usecase.repo.Create(e)
}

func (usecase *ActiveUseCase) Update(id int, e *entity.Active) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}
	
	return usecase.repo.Update(id, e)
}

func (usecase *ActiveUseCase) Delete(id int) error {
	return usecase.repo.Delete(id)
}

func (usecase *ActiveUseCase) FindByID(id int) (task *entity.Active, err error) {
	return usecase.repo.FindByID(id)
}

func (usecase *ActiveUseCase) FindAll() (tasks []*entity.Active, err error) {
	return usecase.repo.FindAll()
}
