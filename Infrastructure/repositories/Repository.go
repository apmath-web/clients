package repositories

import (
	"errors"
	"fmt"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
	"github.com/apmath-web/clients/Infrastructure/mapper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type clientRepository struct {
	db *sqlx.DB
}

func (r *clientRepository) GetClient(id int) (Domain.ClientDomainModelInterface, error) {
	client := new(applicationModels.ClientApplicationModel)
	if err := client.GetClient(id, r.db); err != nil {
		return nil, err
	}
	return mapper.ClientApplicationMapper(*client), nil
}

func (r *clientRepository) insertClient(cl applicationModels.ClientApplicationModel) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	if id, err := cl.SaveClient(tx); err != nil {
		return 0, err
	} else {
		if err = tx.Commit(); err != nil {
			return 0, err
		}
		return id, nil
	}
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
		if err != nil {
			log.Fatal(err.Error())
		}
		repo = &clientRepository{db}
	})
	return repo
}
