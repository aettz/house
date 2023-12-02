package pets

import (
	"fmt"
)

type Pets struct {
	Type  string
	Breed string
	Age   int
	Name  string
}

func (f Pets) Print() {
	fmt.Print("\n")
	fmt.Print("Тип: ", f.Type, "\nПорода: ", f.Breed, "\nИмя: ", f.Name, "\nВозраст: ", f.Age, "\n")

}
