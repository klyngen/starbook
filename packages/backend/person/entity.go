package person

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Name    string
	Picture string
}

func (p *Person) IsValidPerson() bool {
	return len(p.Name) > 0
}

func (p Person) Create(name string, picture string) *Person {
	person := Person{
		Name:    name,
		Picture: name,
	}

	person.CreatedAt = time.Now()

	return &person
}
