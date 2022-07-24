package animal

import "fmt"

type Bear struct {
	_ID   int32
	_name string
	_sexe byte
}

func (d *Bear) Display() {
	fmt.Println("ID:", d.ID())
	fmt.Println("Name:", d.Name())
	fmt.Printf("Sexe: %c\n", d.Sexe())
	fmt.Println("Espece:", d.Espece())
}

func (d *Bear) ID() int32 {
	return d._ID
}

func (d *Bear) Name() string {
	return d._name
}

func (d *Bear) Sexe() byte {
	return d._sexe
}

func (d *Bear) Espece() string {
	return "Bear"
}

func (d *Bear) SetID(id int32) {
	d._ID = id
}

func (d *Bear) SetName(name string) {
	d._name = name
}

func (d *Bear) SetSexe(s byte) {
	d._sexe = s
}
