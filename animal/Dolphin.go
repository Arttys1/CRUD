package animal

import "fmt"

type Dolphin struct {
	_ID   int32
	_name string
	_sexe byte
}

func (d *Dolphin) Display() {
	fmt.Println("ID:", d.ID())
	fmt.Println("Name:", d.Name())
	fmt.Printf("Sexe: %c\n", d.Sexe())
	fmt.Println("Espece:", d.Espece())
}

func (d *Dolphin) ID() int32 {
	return d._ID
}

func (d *Dolphin) Name() string {
	return d._name
}

func (d *Dolphin) Sexe() byte {
	return d._sexe
}

func (d *Dolphin) Espece() string {
	return "Dolphin"
}

func (d *Dolphin) SetID(id int32) {
	d._ID = id
}

func (d *Dolphin) SetName(name string) {
	d._name = name
}

func (d *Dolphin) SetSexe(s byte) {
	d._sexe = s
}
