package repositories

import (
	"fmt"
	"github.com/apmath-web/clients/Domain"
	"github.com/apmath-web/clients/Infrastructure/applicationModels"
	"github.com/apmath-web/clients/Infrastructure/mapper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	if id, err := r.insertClient(client); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func (r *clientRepository) ChangeClient(id int, model Domain.ClientDomainModelInterface) error {
	client := new(applicationModels.ClientApplicationModel)
	if err := client.GetClient(id, r.db); err != nil {
		return err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if err := client.UpdateClient(model, tx); err != nil {
		return err
	}
	return tx.Commit()
}

var repo *clientRepository
var once sync.Once

func dbConfig() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)
}

func GenClientRepository() Domain.ClientRepositoryInterface {
	once.Do(func() {
		db, err := sqlx.Connect("postgres", dbConfig())
		if err != nil {
			log.Fatal(err.Error())
		}
		repo = &clientRepository{db}
	})
	return repo
}
