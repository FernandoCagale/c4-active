package active

import "github.com/FernandoCagale/c4-active/pkg/entity"

type UseCase interface {
	Create(task *entity.Active) (err error)
	Update(id int, active *entity.Active) (err error)
	Delete(id int) (err error)
	FindByID(id int) (active *entity.Active, err error)
	FindAll() (actives []*entity.Active, err error)
}
