package application

import "recu/src/persona/domain"

type CountGenderUc struct {
	db domain.IpersonRepository
}

func NewCountGenderUc(db domain.IpersonRepository)*CountGenderUc{
	return &CountGenderUc{db: db}
}

func (useCase *CountGenderUc)Execute(sexo bool)(int,error){
	return useCase.db.CountGender(sexo)
}