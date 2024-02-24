package pkg

type KotofeyFactory struct{}

func (kf KotofeyFactory) GetJacket() Jacket {
	return KotofeyJacket{}
}

func (kf KotofeyFactory) GetPant() Pant {
	return KotofeyPant{}
}
