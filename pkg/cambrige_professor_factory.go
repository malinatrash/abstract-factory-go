package pkg

type CambridgeProfessorFactory struct{}

func (kf CambridgeProfessorFactory) GetJacket() Jacket {
	return CambridgeProfessorJacket{}
}

func (kf CambridgeProfessorFactory) GetPant() Pant {
	return CambridgeProfessorPant{}
}
