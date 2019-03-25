package models

import "github.com/apmath-web/clients/Domain"

type IdDomainModel struct {
	id int
}

func (i *IdDomainModel) Get() int {
	return i.id
}

func (i *IdDomainModel) Set(id int) {
	i.id = id
}

func GenId(id int) Domain.IdInterface {
	idModel := new(IdDomainModel)
	idModel.id = id
	return idModel
}
