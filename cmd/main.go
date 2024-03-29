package main

import (
	"abstract-factory-go/pkg"
	"fmt"
)

var brands = []string{pkg.Kotofey, pkg.CambridgeProfessor, "Сапожок"}

func main() {

	var factory pkg.Factory

	for factory == nil {
		var choice int
		fmt.Println("---------------------–––––")
		fmt.Println("Выберите линейку костюмов:")
		for index, brand := range brands {
			fmt.Printf("%d. %s\n", index+1, brand)
		}
		fmt.Print(">> ")
		fmt.Scanln(&choice)
		if choice <= len(brands) {
			var err error
			factory, err = pkg.GetFactory(brands[choice-1])
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		} else {
			fmt.Println("Ошибка")
		}
	}

	var batchSize int
	for batchSize <= 0 {
		fmt.Print("Введите размер партии: ")
		fmt.Scanln(&batchSize)
		if batchSize <= 0 {
			fmt.Println("Размер партии должен быть натуральным числом!")
		}
	}

	jacket := factory.GetJacket()
	pant := factory.GetPant()

	fmt.Printf(
		"Был осуществлен заказ на %d %s и %d %s.\n",
		batchSize, jacket.PrintDetails(), batchSize, pant.PrintDetails(),
	)
}
