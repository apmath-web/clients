package viewModels

import (
	"encoding/json"
)

type PassportViewModel struct {
	Series int `json:"series"`
	Number int `json:"number"`
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
