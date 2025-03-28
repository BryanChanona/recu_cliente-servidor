package domain

type IpersonRepository interface {
	SavePerson(Person) error
}