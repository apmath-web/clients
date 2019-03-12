package viewModels

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
