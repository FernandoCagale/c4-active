package active

import (
	"github.com/FernandoCagale/c4-active/config"
	"github.com/FernandoCagale/c4-active/pkg/entity"
)

//GormRepository in memory repo
type GormRepository struct {
	env *config.Config
}

//NewGormRepository create new repository
func NewGormRepository(env *config.Config) *GormRepository {
	return &GormRepository{env}
}

//FindAll all
func (r *GormRepository) FindAll(query *entity.Query) (values []*entity.Active, count int, err error) {
	return []*entity.Active{
		&entity.Active{
			Code: "AMZN",
			Name: "Amazon",
		},
		&entity.Active{
			Code: "FB",
			Name: "Facebook",
		},
	}, 10, nil
}
