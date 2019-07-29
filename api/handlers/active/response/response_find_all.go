package response

import "github.com/FernandoCagale/c4-active/pkg/entity"

//FindAllResponse dto
type FindAllResponse struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

//Data options
type Data struct {
	Count int                `json:"count"`
	Rows  []*FindAllResponse `json:"rows"`
}

//NewResponseFinAll findAll
func NewResponseFinAll(count int, rows []*entity.Active) *Data {
	return &Data{count, convertResponseFinAllCanal(rows)}
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
