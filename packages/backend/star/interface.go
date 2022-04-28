package star

type Writer interface {
	// Create creates a new staring
	Create(star *Star) (*Star, error)
}

type Reader interface {
	// GetRecentStarings returns the n-`amount` of starings
	GetRecentStarings(amount int) ([]Star, error)
	// GetStarsByPerson returns all the stars for a person
	GetStarsByPerson(userId uint) ([]Star, error)
}

type StarRepository interface {
	Writer
	Reader
}
