package family

import "fmt"

type Family struct {
	Sex          string
	Age          int
	Name         string
	FamilyStatus string
	Children     int
}

func (f Family) Print() {
	fmt.Print("\nИмя: ", f.Name, "\nВозраст: ", f.Age, "\nПол: ", f.Sex, "\nСемейный статус: ", f.FamilyStatus, "\nДети: ", f.Children)
}
