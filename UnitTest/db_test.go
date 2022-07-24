package UnitTest

import (
	"CRUD/db"
	"testing"
)

func TestDBInit(t *testing.T) {
	var database db.DB
	database.Init()
	database.Ping()
	database.Close()
}
