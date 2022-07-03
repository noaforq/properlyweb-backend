package entities

type Persons []Person

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Gender    int
	Birthday  string
	Age       int
}
