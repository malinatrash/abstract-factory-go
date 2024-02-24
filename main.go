package main

import "fmt"

// Абстрактные типы костюмов
type Suit interface {
    MakeSuit() string
}

type Jacket interface {
    MakeJacket() string
}

type Trousers interface {
    MakeTrousers() string
}

// Фабрика для производства костюмов
type SuitFactory interface {
    CreateSuit() Suit
}

// Фабрика для производства костюмов линейки "Котофей"
type KotofeyFactory struct{}

func (kf KotofeyFactory) CreateSuit() Suit {
    return KotofeySuit{}
}

// Костюмы линейки "Котофей"
type KotofeySuit struct{}

func (ks KotofeySuit) MakeSuit() string {
    return "Котофей"
}

func (ks KotofeySuit) MakeJacket() string {
    return "Котофей пиджак"
}

func (ks KotofeySuit) MakeTrousers() string {
    return "Котофей брюки"
}

// Фабрика для производства костюмов линейки "Кембриджский профессор"
type ProfessorFactory struct{}

func (pf ProfessorFactory) CreateSuit() Suit {
    return ProfessorSuit{}
}

// Костюмы линейки "Кембриджский профессор"
type ProfessorSuit struct{}

func (ps ProfessorSuit) MakeSuit() string {
    return "Кембриджский профессор"
}

func (ps ProfessorSuit) MakeJacket() string {
    return "Кембриджский профессор пиджак"
}

func (ps ProfessorSuit) MakeTrousers() string {
    return "Кембриджский профессор брюки"
}

// Функция для заказа костюмов определенной линейки и размера партии
func OrderSuits(factory SuitFactory, batchSize int) {
    suit := factory.CreateSuit()
    fmt.Printf("Был осуществлен заказ на %d пиджаков линейки \"%s\" и %d брюк линейки \"%s\".\n",
        batchSize, suit.MakeSuit(), batchSize, suit.MakeSuit())
}

func main() {
    var choice int
    fmt.Println("Выберите линейку костюмов:")
    fmt.Println("1. Котофей")
    fmt.Println("2. Кембриджский профессор")
    fmt.Print(">> ")
    fmt.Scanln(&choice)

    var factory SuitFactory

    switch choice {
    case 1:
        factory = KotofeyFactory{}
    case 2:
        factory = ProfessorFactory{}
    default:
        fmt.Println("Некорректный выбор.")
        return
    }

    var batchSize int
    fmt.Print("Введите размер партии: ")
    fmt.Scanln(&batchSize)

    OrderSuits(factory, batchSize)
}
