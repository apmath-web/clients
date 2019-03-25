package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/clients/Application/v1/validation"
	"github.com/apmath-web/clients/Domain"
)

type PassportViewModel struct {
	Series     int `json:"series"`
	Number     int `json:"number"`
	validation validation.Validation
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

func (p *PassportViewModel) validateNumber() {
	if p.Number < 100000 || p.Number >= 1000000 {
		p.validation.AddMessage(validation.GenMessage("passportNumber", "Isn't right length"))
	}
}

func (p *PassportViewModel) validateSeries() {
	if p.Series < 1000 || p.Series >= 10000 {
		p.validation.AddMessage(validation.GenMessage("passportSeries", "Isn't right length"))
	}
}

func (p *PassportViewModel) Validate() bool {
	p.validateNumber()
	p.validateSeries()
	return p.validation.Empty()
}

func (p *PassportViewModel) GetValidation() Domain.ValidationInterface {
	return &p.validation
}

func (p *PassportViewModel) Hydrate(passport Domain.PassportDomainModelInterface) {
	p.Number = passport.GetNumber()
	p.Series = passport.GetSeries()
}
