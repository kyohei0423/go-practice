package animals

import (
	"testing"
)

func TestDog(t *testing.T) {
	expect := "Dog"
	actual := Dog()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
