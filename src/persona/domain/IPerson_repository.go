package domain

type IpersonRepository interface {
	SavePerson(Person) error
	GetnewPersonIsAdded() (bool, error)
	CountGender(bool)(int,error)
}