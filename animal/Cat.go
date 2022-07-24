package animal

import "fmt"

type Cat struct {
	_ID   int32
	_name string
	_sexe byte
}

func (c *Cat) Display() {
	fmt.Println("ID:", c.ID())
	fmt.Println("Name:", c.Name())
	fmt.Printf("Sexe: %c\n", c.Sexe())
	fmt.Println("Espece:", c.Espece())
}

func (c *Cat) ID() int32 {
	return c._ID
}

func (c *Cat) Name() string {
	return c._name
}

func (c *Cat) Sexe() byte {
	return c._sexe
}

func (c *Cat) Espece() string {
	return "Cat"
}

func (c *Cat) SetID(id int32) {
	c._ID = id
}

func (c *Cat) SetName(name string) {
	c._name = name
}

func (c *Cat) SetSexe(s byte) {
	c._sexe = s
}
