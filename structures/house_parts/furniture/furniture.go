package furniture

import (
	"fmt"
	"hause/structures/house_parts/types"
)

type Furniture struct {
	Type   string
	Length types.Centimeter
	Width  types.Centimeter
	Height types.Centimeter
	Color  string
}

func (f Furniture) Print() {
	fmt.Print("\n")
	fmt.Print("Тип мебели: ", f.Type, "\nДлина: ", f.Length, "\nШирина: ", f.Width, "\nВысота: ", f.Height, "\nЦвет: ", f.Color, "\n")
}
