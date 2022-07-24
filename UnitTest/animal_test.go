package UnitTest

import (
	"CRUD/animal"
	"testing"
)

func TestAnimals(t *testing.T) {
	var (
		cat     animal.Cat
		dog     animal.Dog
		dolphin animal.Dolphin
		bear    animal.Bear
	)

	cat.SetID(1)
	cat.SetName("matou")
	cat.SetSexe('F')

	dog.SetID(2)
	dog.SetName("doggo")
	dog.SetSexe('M')

	dolphin.SetID(3)
	dolphin.SetName("winter")
	dolphin.SetSexe('F')

	bear.SetID(4)
	bear.SetName("ronflex")
	bear.SetSexe('M')

	if cat.ID() != 1 || cat.Name() != "matou" || cat.Sexe() != 'F' || cat.Espece() != "Cat" {
		t.Errorf("cat does not correspond to the expectations")
	}

	if dog.ID() != 2 || dog.Name() != "doggo" || dog.Sexe() != 'M' || dog.Espece() != "Dog" {
		t.Errorf("dog does not correspond to the expectations")
	}

	if dolphin.ID() != 3 || dolphin.Name() != "winter" || dolphin.Sexe() != 'F' || dolphin.Espece() != "Dolphin" {
		t.Errorf("dolphin does not correspond to the expectations")
	}

	if bear.ID() != 4 || bear.Name() != "ronflex" || bear.Sexe() != 'M' || bear.Espece() != "Bear" {
		t.Errorf("bear does not correspond to the expectations")
	}
}
