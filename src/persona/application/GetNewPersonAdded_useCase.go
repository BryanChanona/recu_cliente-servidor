package application

import "recu/src/persona/domain"

type GetNewPersonIsAddedUc struct {
	db domain.IpersonRepository
}

func NewGetNewPersonIsAddedUc(db domain.IpersonRepository)*GetNewPersonIsAddedUc{
	return &GetNewPersonIsAddedUc{db: db}
}

func (useCase *GetNewPersonIsAddedUc)Execute()(bool, error){
	return useCase.db.GetnewPersonIsAdded()
}