package repo

// ----------------------------------------------------------------------------
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// ----------------------------------------------------------------------------
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}
