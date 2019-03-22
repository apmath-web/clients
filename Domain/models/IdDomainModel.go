package models

type Id struct {
	id int
}

func (i *Id) Get() int {
	return i.id
}

func (i *Id) Set(id int) {
	i.id = id
}
