package applicationModels

import (
	"database/sql"
	"github.com/apmath-web/clients/Domain"
	"github.com/jmoiron/sqlx"
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

func (p *PassportApplicationModel) GetPassport(clientId int, db *sqlx.DB) error {
	pas := PassportApplicationModel{}
	if err := db.Get(&pas, "SELECT * FROM passports WHERE client_id=$1", clientId); err != nil {
		return err
	}
	p.Id = pas.Id
	p.ClientId = pas.ClientId
	p.Series = pas.Series
	p.Number = pas.Number
	return nil
}

func (p *PassportApplicationModel) UpdatePassport(pas Domain.PassportDomainModelInterface, tx *sql.Tx) error {
	if p.Series != pas.GetSeries() {
		if _, err := tx.Exec("UPDATE passports SET series=$2 WHERE id = $1",
			p.Id, pas.GetSeries()); err != nil {
			return err
		}
	}
	if p.Number != pas.GetNumber() {
		if _, err := tx.Exec("UPDATE passports SET number=$2 WHERE id = $1",
			p.Id, pas.GetNumber()); err != nil {
			return err
		}
	}
	return nil
}
