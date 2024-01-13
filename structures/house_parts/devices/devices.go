package devices

import (
	"fmt"
	"hause/structures/house_parts/types"
)

type Devices struct {
	Type     string
	Length   types.Centimeter
	Width    types.Centimeter
	Color    string
	Wireless bool
}

func (d Devices) Print() {
	fmt.Print("\nТип устройства: ", d.Type, "\nДлина устройства: ", d.Length, "\nШирина устройства: ", d.Width, "\nЦвет: ", d.Color, "\n")
	if d.Wireless {
		fmt.Print("Поддерживает беспроводную работу ", "\n")
	} else {
		fmt.Print("Не поддерживает беспроводную работу ", "\n")
	}
}
