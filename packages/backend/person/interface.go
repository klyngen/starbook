package person

type PersonRepository interface {
	// Create creates a person
	Create(person *Person) (*Person, error)
	AllPersons() ([]Person, error)
}
