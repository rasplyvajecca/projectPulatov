package house

import (
	"fmt"
	"projectPulatov/projectPulatov/family"
	"projectPulatov/projectPulatov/furniture"
	"projectPulatov/projectPulatov/pets"
	"projectPulatov/projectPulatov/plants"
	"projectPulatov/projectPulatov/relatives"
)

type House struct {
	Rooms          int
	Floors         int
	Length         float64
	Width          float64
	Height         float64
	HouseFamily    []family.Family
	HouseRelatives []relatives.Relatives
	HousePets      []pets.Pets
	HouseFurniture []furniture.Furniture
	HousePlants    []plants.Plants
}

func NewHouse() House {
	return House{
		Rooms:          3,
		Floors:         2,
		Length:         39.6,
		Width:          45.2,
		Height:         3,
		HouseFamily:    family.NewMember(),
		HouseRelatives: relatives.NewRelative(),
		HousePets:      pets.NewPet(),
		HouseFurniture: furniture.NewObject(),
		HousePlants:    plants.NewPlant(),
	}
}

func PrintHouse() {
	house := NewHouse()
	fmt.Printf("Новый дом:\n")
	fmt.Printf("Кол-во комнат: %d\n", house.Rooms)
	fmt.Printf("Кол-во этажей: %d кв. м\n", house.Floors)
	fmt.Printf("Длина: %.1f м\n", house.Length)
	fmt.Printf("Ширина: %.1f м\n", house.Width)
	fmt.Printf("Высота: %.1f м\n", house.Height)

	fmt.Println("Члены семьи:")
	for _, object := range house.HouseFamily {
		fmt.Printf("- %s: возраст: %d\n", object.Member, object.Age)
	}

	fmt.Println("Родственники:")
	for _, object := range house.HouseRelatives {
		fmt.Printf("- %s: возраст: %d\n", object.Member, object.Age)
	}

	fmt.Println("Домашние животные:")
	for _, object := range house.HousePets {
		fmt.Printf("- %s: %d шт.\n", object.Pet, object.Count)
	}

	fmt.Println("Мебель:")
	for _, object := range house.HouseFurniture {
		fmt.Printf("- %s: %d шт., %.2f кв. м\n", object.Name, object.Count, object.Size)
	}

	fmt.Println("Растения:")
	for _, object := range house.HousePlants {
		fmt.Printf("- %s: %d шт.\n", object.Plant, object.Count)
	}
}
