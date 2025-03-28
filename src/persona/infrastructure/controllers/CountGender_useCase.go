package controllers

import "recu/src/persona/application"

type CountGenderUc struct {
	useCase *application.CountGenderUc
}

func NewCountGenderUc(useCase *application.CountGenderUc)*CountGenderUc{
 return &CountGenderUc{useCase: useCase}
}

func (controller *CountGenderUc) Execute()(int, error){
	
}