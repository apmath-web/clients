package applicationModels

import "github.com/apmath-web/clients/Domain"

type PassportApplicationModel struct {
	Series int
	Number int
}

func (p *PassportApplicationModel) GetSeries() int {
	return p.Series
}

func (p *PassportApplicationModel) GetNumber() int {
	return p.Number
}

func (p *PassportApplicationModel) Hydrate(passport Domain.PassportDomainModelInterface) {
	p.Number = passport.GetNumber()
	p.Series = passport.GetSeries()
}
