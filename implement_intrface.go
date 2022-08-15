package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	ani := getAnimal(num)

	ani.Eat()
	ani.Drink()
	ani.Sleep()

	c := &cat{}
	c.propertyName()
	c.attributeMeaning()
}

func getAnimal(num int) animal {
	if num%2 == 0 {
		return &cat{}
	} else {
		return &dog{}
	}
}

type animal interface {
	Eat()
	Drink()
	Sleep()
}

type cat struct {
}

func (c *cat) Eat() {
	fmt.Println("Cat eat fish.")
}

func (c *cat) Drink() {
	fmt.Println("Cat drink water.")
}

func (c *cat) Sleep() {
	fmt.Println("Cat sleep 6 hours.")
}

type dog struct {
}

func (d *dog) Eat() {
	fmt.Println("Dog eat bone.")
}

func (d *dog) Drink() {
	fmt.Println("Dog drink milk.")
}

func (d *dog) Sleep() {
	fmt.Println("Dog sleep hours.")
}

type property interface {
	propertyName()
	attributeMeaning()
}

func (c *cat) propertyName() {
	fmt.Println("Cat.")
}

func (c *cat) attributeMeaning() {
	fmt.Println("Cat proud.")
}

func (d *dog) propertyName() {
	fmt.Println("Dog")
}

func (d *dog) attributeMeaning() {
	fmt.Println("Dog loyalty.")
}
