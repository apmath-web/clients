package models

import "github.com/apmath-web/clients/Domain"

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

func GenPassportDomainModel(series int, number int) Domain.PassportDomainModelInterface {
	dm := new(PassportDomainModel)
	dm.Number = number
	dm.Series = series
	return dm
}
