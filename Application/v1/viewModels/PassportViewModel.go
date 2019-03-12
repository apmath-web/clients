package viewModels

import "encoding/json"

type PassportViewModel struct {
	series int
	number int
}

func (p *PassportViewModel) GetSeries() int {
	return p.series
}

func (p *PassportViewModel) GetNumber() int {
	return p.number
}

func (p *PassportViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]int{
		"series": p.series,
		"number": p.number,
	})
}
