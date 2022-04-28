package star

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/klyngen/starbook/person"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createTestDatabase() GormStarRepository {
	connection, _ := gorm.Open(sqlite.Open("test.db"))
	connection.AutoMigrate(&person.Person{})
	connection.AutoMigrate(&Star{})

	connection.Create(&person.Person{
		Name: "Not Klingen",
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
		},
	})
	connection.Create(&person.Person{
		Name: "Klingen",
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
		},
	})

	return GormStarRepository{
		db: connection,
	}
}

func TestCreateStaring(t *testing.T) {
	defer os.Remove("test.db")

	repo := createTestDatabase()

	star, err := repo.Create(&Star{
		Comment: "Testing purpouses",
		Sender: person.Person{
			Model: gorm.Model{
				ID: 1,
			},
		},
		Recipient: person.Person{
			Model: gorm.Model{
				ID: 2,
			},
		},
	})

	if err != nil {
		t.Log(err)
		t.Log("Tried to create a staring but it failed")
		t.Fail()
	}

	if star.ID < 1 {
		t.Fail()
	}

}
func TestGetStarsForUser(t *testing.T) {
	defer os.Remove("test.db")

	repo := createTestDatabase()

	for i := 0; i < 25; i++ {
		repo.Create(&Star{
			Comment:     fmt.Sprintf("Testing %v", i),
			SenderID:    uint(1),
			RecipientID: uint(2),
		})
	}

	stars, err := repo.GetStarsByPerson(2)

	if err != nil {
		t.Log(err)
		t.Log("Tried to get starings for a user but it failed")
		t.Fail()
	}

	if len(stars) != 25 {
		t.Logf("Amount of stars are wrong. Expected 25 got %v", len(stars))
		t.Fail()
	}
	assert.Equal(t, "Not Klingen", stars[0].Sender.Name)
	assert.Equal(t, "Klingen", stars[0].Recipient.Name)
}

func TestGetRecentStarings(t *testing.T) {
	defer os.Remove("test.db")
	repo := createTestDatabase()

	for i := 0; i < 25; i++ {
		repo.Create(&Star{
			Comment:     fmt.Sprintf("Testing %v", i),
			SenderID:    uint(1),
			RecipientID: uint(2),
		})
	}

	stars, err := repo.GetRecentStarings(10)

	assert.NoError(t, err)
	assert.Len(t, stars, 10)

}
