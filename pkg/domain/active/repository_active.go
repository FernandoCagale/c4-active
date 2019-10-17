package active

import (
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"github.com/FernandoCagale/c4-active/pkg/infra/errors"
	"strconv"
)

type ActiveRepository struct {
	m map[string]*entity.Active
}

func NewRepository() *ActiveRepository {
	amazon := &entity.Active{
		ID:   1,
		Code: "AMZN",
		Name: "Amazon",
	}

	facebook := &entity.Active{
		ID:   2,
		Code: "FB",
		Name: "Facebook",
	}

	var m = map[string]*entity.Active{"1": amazon, "2": facebook}
	return &ActiveRepository{m}
}

func (repo *ActiveRepository) Create(e *entity.Active) error {
	if repo.m[strconv.Itoa(e.ID)] != nil {
		return errors.ErrInvalidPayload
	}
	repo.m[strconv.Itoa(e.ID)] = e
	return nil
}

func (repo *ActiveRepository) Update(id int, e *entity.Active) error {
	if repo.m[strconv.Itoa(id)] == nil {
		return errors.ErrNotFound
	}
	repo.m[strconv.Itoa(id)] = e
	return nil
}

func (repo *ActiveRepository) Delete(id int) error {
	if repo.m[strconv.Itoa(id)] == nil {
		return errors.ErrNotFound
	}
	delete(repo.m, strconv.Itoa(id))
	return nil
}

func (repo *ActiveRepository) FindByID(id int) (task *entity.Active, err error) {
	if repo.m[strconv.Itoa(id)] == nil {
		return nil, errors.ErrNotFound
	}
	return repo.m[strconv.Itoa(id)], nil
}

func (repo *ActiveRepository) FindAll() (actives []*entity.Active, err error) {
	for _, active := range repo.m {
		actives = append(actives, active)
	}
	return actives, nil
}
