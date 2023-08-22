package model

import "strconv"

var (
	names []string
)

func init() {
	names = []string{"buick", "audi", "ford"}

}

type Car struct {
	Name     string
	CarDoor  *Door
	CarWheel *Wheel
}

type Door struct {
	Age  int
	Name string
}

type Wheel struct {
	Count int
	Name  string
}

func NewCar(idx int) *Car {
	c := Car{}
	c.Name = names[idx%3]
	c.CarDoor = &Door{Age: idx, Name: "door:" + strconv.Itoa(idx)}
	c.CarWheel = &Wheel{Count: idx, Name: "wheel:" + strconv.Itoa(idx)}
	return &c
}
