package star

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStarRepository struct {
	db *gorm.DB
}

func NewGormStarRepository(host string, user string, password string, database string) (*GormStarRepository, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Europe/Oslo", host, user, password, database)
	connection, err := gorm.Open(postgres.Open(connectionString))

	connection.AutoMigrate(&Star{})

	return &GormStarRepository{
		db: connection,
	}, err
}

func (r *GormStarRepository) Create(star *Star) (*Star, error) {
	result := r.db.Create(star)
	return star, result.Error
}

func (r *GormStarRepository) GetRecentStarings(amount int) ([]Star, error) {
	var starings []Star
	res := r.db.Order("created_at desc").Limit(amount).Find(&starings)
	return starings, res.Error
}

func (r *GormStarRepository) GetStarsByPerson(userId uint) ([]Star, error) {
	var starings []Star
	log.Println(userId)
	res := r.db.Joins("Sender").Joins("Recipient").Where("recipient_id", userId).Order("stars.created_at desc").Find(&starings)
	return starings, res.Error
}
