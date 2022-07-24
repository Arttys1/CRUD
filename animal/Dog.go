package animal

import "fmt"

type Dog struct {
	_ID   int32
	_name string
	_sexe byte
}

func (d *Dog) Display() {
	fmt.Println("ID:", d.ID())
	fmt.Println("Name:", d.Name())
	fmt.Printf("Sexe: %c\n", d.Sexe())
	fmt.Println("Espece:", d.Espece())
}

func (d *Dog) ID() int32 {
	return d._ID
}

func (d *Dog) Name() string {
	return d._name
}

func (d *Dog) Sexe() byte {
	return d._sexe
}

func (d *Dog) Espece() string {
	return "Dog"
}

func (d *Dog) SetID(id int32) {
	d._ID = id
}

func (d *Dog) SetName(name string) {
	d._name = name
}

func (d *Dog) SetSexe(s byte) {
	d._sexe = s
}
