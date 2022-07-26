/**
 * Package to manage Animal in the database.
 * This package is used to create, read, update and delete the database.
 */
package animal

import (
	"CRUD/db"
	"database/sql"
	"fmt"
)

type Animal interface {
	ID() int32
	Name() string
	Sexe() byte
	Espece() string
	SetID(int32)
	SetName(string)
	SetSexe(byte)
	Display()
}

func UpdateAnimalInDB(database *db.DB, animal Animal) {
	query := "UPDATE animal SET "
	query += fmt.Sprintf("nom = '%s', ", animal.Name())
	query += fmt.Sprintf("sexe = '%c' ", animal.Sexe())
	query += fmt.Sprintf("WHERE id = %d;", animal.ID())
	_, err := database.Exec(query)
	if err != nil {
		panic(query + "\n" + err.Error())
	}
}

func DeleteAnimalInDB(database *db.DB, id int32) {
	query := "DELETE FROM animal WHERE id = " + fmt.Sprintf("%d;", id)
	_, err := database.Exec(query)
	if err != nil {
		panic(query + "\n" + err.Error())
	}
}

func AddAnimalInDB(database *db.DB, animal Animal) {
	query := "INSERT INTO animal (id, nom, sexe, espece_id) VALUES ("
	query += fmt.Sprintf("%d, ", animal.ID())
	query += fmt.Sprintf("'%s', ", animal.Name())
	query += fmt.Sprintf("'%c', ", animal.Sexe())
	query += fmt.Sprintf("(SELECT id FROM espece WHERE espece = '%s'));", animal.Espece())

	_, err := database.Exec(query)
	if err != nil {
		panic(query + "\n" + err.Error())
	}
}

func GetMaxIDFromDB(database *db.DB) int32 {
	rows, err := database.Query("SELECT MAX(id) FROM animal;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var maxID int32
	rows.Scan(&maxID)
	return maxID
}

func GetAnimal(database *db.DB, id int32) Animal {
	rows, err := database.Query("SELECT animal.id, nom, sexe, espece " +
		"FROM animal JOIN espece " +
		"ON animal.espece_id = espece.id " +
		"WHERE animal.id = " + fmt.Sprintf("%d;", id))

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var animal Animal
	if rows.Next() {
		animal = loadAnimal(rows)
	} else {
		panic(fmt.Sprintf("Animal with id %d not found", id))
	}

	return animal
}

func GetAllAnimals(database *db.DB) []Animal {
	rows, err := database.Query("SELECT animal.id, nom, sexe, espece " +
		"FROM animal JOIN espece " +
		"ON animal.espece_id = espece.id order by id;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		animal := loadAnimal(rows)
		animals = append(animals, animal)
	}

	return animals
}

func loadAnimal(rows *sql.Rows) Animal {
	var (
		id     int32
		name   string
		sexe   []byte
		espece string
		animal Animal
	)
	err := rows.Scan(&id, &name, &sexe, &espece)
	if err != nil {
		panic(err)
	}
	switch espece {
	case "cat":
		animal = &Cat{
			_ID:   id,
			_name: name,
			_sexe: sexe[0],
		}
	case "dog":
		animal = &Dog{
			_ID:   id,
			_name: name,
			_sexe: sexe[0],
		}
	case "bear":
		animal = &Bear{
			_ID:   id,
			_name: name,
			_sexe: sexe[0],
		}
	case "dolphin":
		animal = &Dolphin{
			_ID:   id,
			_name: name,
			_sexe: sexe[0],
		}
	default:
		panic(fmt.Sprintf("Unknown espece. id : %d", id))
	}

	return animal
}

func DisplayAll(animals []Animal) {
	for _, animal := range animals {
		fmt.Println("-----------")
		animal.Display()
	}
}
