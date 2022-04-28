package person

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonGormRepository struct {
	db *gorm.DB
}

func NewGormPersonRepository(host string, user string, password string, database string) (*PersonGormRepository, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Europe/Oslo", host, user, password, database)
	connection, err := gorm.Open(postgres.Open(connectionString))

	connection.AutoMigrate(&Person{})

	return &PersonGormRepository{
		db: connection,
	}, err
}

func (p *PersonGormRepository) Create(person *Person) (*Person, error) {
	res := p.db.Create(person)
	return person, res.Error
}

func (p *PersonGormRepository) AllPersons() ([]Person, error) {
	var persons []Person

	err := p.db.Find(&persons)

	return persons, err.Error
}
