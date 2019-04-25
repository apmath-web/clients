package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type clientRepository struct {
	db *sqlx.DB
}

func (r *clientRepository) GetClient(id int) (Domain.ClientDomainModelInterface, error) {
	return nil, errors.New("No client with such id")
}

func (r *clientRepository) insertClient(cl applicationModels.ClientApplicationModel) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var clientId int
	err = tx.QueryRow("INSERT INTO clients "+
		"(first_name, last_name, birth_date, sex, marital_status, children) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6) RETURNING id",
		cl.GetFirstName(), cl.GetLastName(), cl.GetBirthDate(),
		cl.GetSex(), cl.GetMaritalStatus(), cl.GetChildren()).Scan(&clientId)
	fmt.Printf("%+v", clientId)
	if err != nil {
		return 0, err
	}
	if _, err = r.InsertPassport(cl.GetPassport(), tx, clientId); err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return clientId, nil
}

func (r *clientRepository) InsertPassport(pas applicationModels.PassportApplicationModel, tx *sql.Tx, clientId int) (sql.Result, error) {
	return tx.Exec("INSERT INTO passports "+
		"(series, number, client_id) VALUES "+
		"($1, $2, $3) RETURNING id",
		pas.GetSeries(), pas.GetNumber(), clientId)
}

func (r *clientRepository) SetClient(model Domain.ClientDomainModelInterface) (int, error) {
	var client applicationModels.ClientApplicationModel
	client.Hydrate(model)
	fmt.Printf("%+v\n", model.GetPassport())
	if id, err := r.insertClient(client); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func (r *clientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	return errors.New("No client with such id")
}

var repo *clientRepository
var once sync.Once

const DB_CONNECT_STRING = "host=localhost port=5432 user=test password=test dbname=test sslmode=disable"

func GenClientRepository() Domain.ClientRepositoryInterface {
	once.Do(func() {
		db, err := sqlx.Connect("postgres", DB_CONNECT_STRING)
		fmt.Printf("%T", db)
		if err != nil {
			log.Fatal(err.Error())
		}
		repo = &clientRepository{db}
	})
	return repo
}
