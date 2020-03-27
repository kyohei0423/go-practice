package animals

import "testing"

func TestCat(t *testing.T) {
	expect := "Cat"
	actual := Cat()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
