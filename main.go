package main

import (
	"CRUD/animal"
	"CRUD/db"
)

func main() {
	var (
		database db.DB
		config   = db.Config{
			User:         "root",
			Password:     "",
			DatabaseName: "bestiole",
			Address:      "127.0.0.1",
			Port:         "3306",
		}
	)

	database.Init(config)
	defer database.Close()

	var animals []animal.Animal = animal.GetAllAnimals(&database)

	animal.DisplayAll(animals)
}
