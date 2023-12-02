package house

import (
	"fmt"
	"hause/structures/house_parts/devices"
	"hause/structures/house_parts/family"
	"hause/structures/house_parts/furniture"
	"hause/structures/house_parts/pets"
	"hause/structures/house_parts/types"
)

type House struct {
	Family        []family.Family
	Devices       []devices.Devices
	Furniture     []furniture.Furniture
	Pets          []pets.Pets
	Rooms         int
	Square        types.SquareMeter
	CeilingHeight types.Centimeter
}

func (h House) Print() {
	fmt.Print("Комнаты: ", h.Rooms, "\nПлощадь: ", h.Square, "\nВысота потолков: ", h.CeilingHeight, "\n")
	for _, families := range h.Family {
		families.Print()
	}
	for _, pets := range h.Pets {
		pets.Print()
	}
	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}
func Make() House {
	table := furniture.Furniture{
		Type:   "Стол",
		Length: 160,
		Width:  70,
		Height: 100,
		Color:  "Серый",
	}
	chair := furniture.Furniture{
		Type:   "Кресло",
		Length: 100,
		Width:  45,
		Height: 70,
		Color:  "Черный",
	}
	bed := furniture.Furniture{
		Type:   "Кровать",
		Length: 200,
		Width:  100,
		Height: 20,
		Color:  "Розовый",
	}
	wardrobe := furniture.Furniture{
		Type:   "Гардероб",
		Length: 400,
		Width:  300,
		Height: 250,
		Color:  "Белый",
	}

	mom := family.Family{
		Sex:          "женщина",
		Age:          51,
		Name:         "Вера",
		FamilyStatus: "замужем",
		Children:     3,
	}
	dad := family.Family{
		Sex:          "мужчина",
		Age:          56,
		Name:         "Олег",
		FamilyStatus: "женат",
		Children:     3,
	}
	bro := family.Family{
		Sex:          "мужчина",
		Age:          20,
		Name:         "Савелий",
		FamilyStatus: "свободен",
		Children:     0,
	}
	esis := family.Family{
		Sex:          "женщина",
		Age:          24,
		Name:         "Александра",
		FamilyStatus: "свободна",
		Children:     0,
	}
	elsis := family.Family{
		Sex:          "женщина",
		Age:          28,
		Name:         "Ксения",
		FamilyStatus: "свободна",
		Children:     0,
	}
	doggy := pets.Pets{
		Type:  "Собака",
		Breed: "Такса",
		Age:   14,
		Name:  "Тася",
	}

	PC := devices.Devices{
		Type:     "PC",
		Length:   50,
		Width:    30,
		Color:    "Белый",
		Wireless: false,
	}
	phone := devices.Devices{
		Type:     "Телефон",
		Length:   16,
		Width:    8,
		Color:    "Зеленый",
		Wireless: true,
	}
	mouse := devices.Devices{
		Type:     "Мышка",
		Length:   15,
		Width:    7,
		Color:    "Бирюзовый",
		Wireless: true,
	}
	headphones := devices.Devices{
		Type:     "наушники",
		Length:   10,
		Width:    1,
		Color:    "Черный",
		Wireless: false,
	}
	keyboard := devices.Devices{
		Type:     "Клавиатура",
		Length:   20,
		Width:    7,
		Color:    "Белый",
		Wireless: true,
	}
	monitor := devices.Devices{
		Type:     "Монитор",
		Length:   35,
		Width:    20,
		Color:    "Черный",
		Wireless: false,
	}
	house := House{
		Family:        []family.Family{dad, mom, bro, esis, elsis},
		Pets:          []pets.Pets{doggy},
		Devices:       []devices.Devices{PC, phone, keyboard, monitor, mouse, headphones},
		Furniture:     []furniture.Furniture{table, chair, wardrobe, bed},
		Rooms:         3,
		Square:        70,
		CeilingHeight: 270,
	}
	return house
}
