package model

import "strconv"

var (
	names []string
)

func init() {
	names = []string{"buick", "audi", "ford"}

}

type Car struct {
	Name     string `json:"name"`
	CarDoor  *Door  `json:"door"`
	CarWheel *Wheel `json:"wheel"`
}

type Door struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type Wheel struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

func NewCar(idx int) *Car {
	c := Car{}
	c.Name = names[idx%3]
	c.CarDoor = &Door{Age: idx, Name: "door:" + strconv.Itoa(idx)}
	c.CarWheel = &Wheel{Count: idx, Name: "wheel:" + strconv.Itoa(idx)}
	return &c
}
