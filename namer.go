// Package namer implements a Heroku-like random name-generator.
//
// Names have the format {noun}-{adjective}-{number}, and are generated
// using a random noun, random adjective, and random four digit number
// between 1000 and 9999. The noun and adjective are chosen from a fixed
// list of 74 nouns and 61 adjectives.
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
	nAdjective = len(adjectives)
	nNoun      = len(nouns)
)

func (n *defaultNamer) Name() string {
	adjective := adjectives[rand.Intn(nAdjective)]
	noun := nouns[rand.Intn(nNoun)]
	digits := rand.Intn(ub-lb) + lb
	return fmt.Sprintf("%s-%s-%d", adjective, noun, digits)
}
