package UnitTest

import (
	"CRUD/animal"
	"CRUD/db"
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

func TestAddAnimalInDB(t *testing.T) {
	var (
		database db.DB
		config   = db.Config{
			User:         "root",
			Password:     "",
			DatabaseName: "bestiole",
			Address:      "127.0.0.1",
			Port:         "3306",
		}
		cat animal.Cat
	)
	cat.SetID(1000)
	cat.SetName("testDB")
	cat.SetSexe('M')

	database.Init(config)
	defer database.Close()
	animal.AddAnimalInDB(&database, &cat)

	rows, err := database.Query("SELECT id, nom, sexe, espece FROM animal natural join espece where id = 1000")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id     int
			name   string
			sexe   string
			espece string
		)
		err = rows.Scan(&id, &name, &sexe, &espece)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if id != 1000 && name != "testDB" && sexe != "M" && espece != "Cat" {
			t.Errorf("cat does not correspond to the expectations")
		}
	}

	//clean DB after test
	animal.DeleteAnimalInDB(&database, 1000)
}
