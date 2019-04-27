package applicationModels

import (
	"database/sql"
	"github.com/apmath-web/clients/Domain"
)

type PassportApplicationModel struct {
	Id       int `db:"id"`
	Series   int `db:"series"`
	Number   int `db:"number"`
	ClientId int `db:"client_id"`
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

func (p *PassportApplicationModel) SavePassport(tx *sql.Tx) error {
	var passportId int
	if err := tx.QueryRow("INSERT INTO passports "+
		"(series, number, client_id) VALUES "+
		"($1, $2, $3) RETURNING id",
		p.GetSeries(), p.GetNumber(), p.ClientId).Scan(&passportId); err != nil {
		return err
	}
	p.Id = passportId
	return nil
}
