package active

import "github.com/FernandoCagale/c4-active/pkg/entity"

//UseCase use case interface
type UseCase interface {
	FindAll(query *entity.Query) (values []*entity.Active, count int, err error)
}
