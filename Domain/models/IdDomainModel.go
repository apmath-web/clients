package models

import "github.com/apmath-web/clients/Domain"

type Id struct {
	id int
}

func (i *Id) Get() int {
	return i.id
}

func (i *Id) Set(id int) {
	i.id = id
}

func GenId(id int) Domain.IdInterface {
	idModel := new(Id)
	idModel.id = id
	return idModel
}
