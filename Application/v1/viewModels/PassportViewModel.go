package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Domain"
)

type PassportViewModel struct {
	Series     int `json:"series"`
	Number     int `json:"number"`
	validation Domain.ValidationInterface
}

func (p *PassportViewModel) GetSeries() int {
	return p.Series
}

func (p *PassportViewModel) GetNumber() int {
	return p.Number
}

func (p *PassportViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]int{
		"Series": p.Series,
		"Number": p.Number,
	})
}

func (p *PassportViewModel) Validate() bool {
	return true
}

func (p *PassportViewModel) GetValidation() Domain.ValidationInterface {
	return p.validation
}
