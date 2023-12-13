package structs

import "fmt"

type customer struct {
	firstname string
	lastname  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi : ", c.firstname)
}
func (c customer) update() {
	fmt.Println("Güncellendi : ", c.firstname)
}
func Demo2() {
	c := customer{firstname: "Özgür", lastname: "Karakuş", age: 22}
	c.save()
	c.update()
}
