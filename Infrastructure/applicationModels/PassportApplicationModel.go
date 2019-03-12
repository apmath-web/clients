package applicationModels

type PassportApplicationModel struct {
	series int
	number int
}

func (p *PassportApplicationModel) GetSeries() int {
	return p.series
}

func (p *PassportApplicationModel) GetNumber() int {
	return p.number
}

func (p *PassportApplicationModel) SetNumber(number int) {
	p.number = number
}
func (p *PassportApplicationModel) SetSeries(series int) {
	p.series = series
}
