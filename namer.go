package namer

import (
	"fmt"
	"math/rand"
)

// Namer specifies an interface for generating random names.
type Namer interface {
	// Name returns a random name.
	Name() string
}

// New creates a new Namer.
func New() Namer {
	return &defaultNamer{}
}

type defaultNamer struct{}

const (
	lb = 1000
	ub = 9999
)

var (
	nNoun      = len(nouns)
	nAdjective = len(adjectives)
)

func (n *defaultNamer) Name() string {
	noun := nouns[rand.Intn(nNoun)]
	adjective := adjectives[rand.Intn(nAdjective)]
	digits := rand.Intn(ub-lb) + lb
	return fmt.Sprintf("%s-%s-%d", noun, adjective, digits)
}
