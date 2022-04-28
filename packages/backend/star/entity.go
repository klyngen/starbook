package star

import (
	"github.com/klyngen/starbook/person"
	"gorm.io/gorm"
)

type Star struct {
	gorm.Model
	Comment     string
	Recipient   person.Person `gorm:"foreignKey:ID;references:recipient_id"`
	RecipientID uint
	Sender      person.Person `gorm:"foreignKey:ID;references:sender_id"`
	SenderID    uint
}

func (s *Star) IsValidStar() bool {
	return len(s.Comment) > 0
}
