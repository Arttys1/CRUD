package main

import (
	"CRUD/animal"
	"CRUD/db"
)

func main() {
	var db = db.DB{}

	defer db.Close()

	db.Init()

	var animals []animal.Animal = animal.GetAllAnimals(&db)

	animal.DisplayAll(animals)
}
