package response

import (
	"github.com/FernandoCagale/c4-active/pkg/entity"
	"os"
)

//FindAllResponse dto
type FindAllResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

//Data options
type Data struct {
	Count   int                `json:"count"`
	Rows    []*FindAllResponse `json:"rows"`
	Version string             `json:"version"`
}

//NewResponseFinAll findAll
func NewResponseFinAll(count int, rows []*entity.Active) *Data {
	version := os.Getenv("VERSION")

	if version ==  "" {
		version = "V1"
	}

	return &Data{
		Count: count,
		Rows:  convertResponseFinAllCanal(rows),
		Version: version,
	}
}

func convertResponseFinAllCanal(actives []*entity.Active) (rows []*FindAllResponse) {
	for _, active := range actives {
		row := &FindAllResponse{
			Code: active.Code,
			Name: active.Name,
		}
		rows = append(rows, row)
	}
	return rows
}
