package animal

import "fmt"

type Name string

type Dog struct {
	Name Name
}

func NewName() Name {
	return Name("Dog")
}

func NewDog(n Name) Dog {
	return Dog{
		Name: n,
	}
}

func (d *Dog) Cry() {
	fmt.Println("wan")
}
