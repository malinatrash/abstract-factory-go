package pkg

import (
	"fmt"
)

type Factory interface {
	GetJacket() Jacket
	GetPant() Pant
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	case Kotofey:
		return &KotofeyFactory{}, nil
	case CambridgeProfessor:
		return &CambridgeProfessorFactory{}, nil
	default:
		return nil, fmt.Errorf("производитель %s - не найден", brand)
	}
}
