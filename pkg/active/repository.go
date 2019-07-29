package active

import "github.com/FernandoCagale/c4-active/pkg/entity"

//Repository repository interface
type Repository interface {
	FindAll(query *entity.Query) (values []*entity.Active, count int, err error)
}
