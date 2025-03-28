package application

import "recu/src/persona/domain"

type AddPersonUc struct {
	db domain.IpersonRepository
}

func NewAddPersonUc(db domain.IpersonRepository)*AddPersonUc{
	return &AddPersonUc{db:db}
}

func (useCase *AddPersonUc)Execute(person domain.Person) error{
	return useCase.db.SavePerson(person)
}