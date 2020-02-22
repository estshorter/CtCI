package animal

import "time"

// Kind determine dog or cat
type Kind int

// defines AnimalType
const (
	Cat Kind = iota
	Dog
)

// Animal defines animal type and added time
type Animal struct {
	Kind
	time.Time
}
