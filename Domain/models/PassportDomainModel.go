package models

type PassportDomainModel struct {
	Series int
	Number int
}

func (p *PassportDomainModel) GetSeries() int {
	return p.Series
}

func (p *PassportDomainModel) GetNumber() int {
	return p.Number
}
