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
