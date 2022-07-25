package UnitTest

import (
	"CRUD/db"
	"testing"
)

func TestDBInit(t *testing.T) {
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
	database.Ping()
	database.Close()
}

func TestQuery(t *testing.T) {
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

	rows, err := database.Query("SELECT id, nom, sexe, espece FROM animal natural join espece")
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
		t.Logf("id: %d, name: %s, sexe: %s, espece: %s", id, name, sexe, espece)
	}
}
