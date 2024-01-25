package namer

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	namer := New()
	_, ok := namer.(Namer)
	if !ok {
		t.Fatalf("New() returned type %T, want type Namer", namer)
	}
}

const (
	adjectiveIdx = 0
	nounIdx      = 1
	digitsIdx    = 2
)

var testIterations = 1000

func TestName(t *testing.T) {
	namer := New()

	for i := 0; i < testIterations; i++ {
		name := namer.Name()

		parts := strings.Split(name, "-")
		if len(parts) != 3 {
			t.Fatalf("Name() returned malformed name: %s", name)
		}

		checkAdjective(t, parts[adjectiveIdx])
		checkNoun(t, parts[nounIdx])
		checkDigits(t, parts[digitsIdx])
	}
}

func checkNoun(t *testing.T, noun string) {
	t.Helper()

	if !slices.Contains(nouns, noun) {
		t.Fatalf("Name() returned name with unexpected noun: %s", noun)
	}
}

func checkAdjective(t *testing.T, adjective string) {
	t.Helper()

	if !slices.Contains(adjectives, adjective) {
		t.Fatalf("Name() returned name with unexpected adjective: %s", adjective)
	}
}

func checkDigits(t *testing.T, digits string) {
	t.Helper()

	d, err := strconv.Atoi(digits)
	if err != nil {
		t.Fatalf("Name() returned name with invalid digits: %s", digits)
	}

	if d < lb || d >= ub {
		t.Fatalf("Name() returned name with out of bound digits: %d", d)
	}
}
