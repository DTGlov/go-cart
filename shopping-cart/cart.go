package main

import (
	"fmt"
	"os"
)


type cart struct {
	name  string
	items map[string]float64
}

//making a new cart
func newCart(name string) cart {
	c := cart{
		name:  name,
		items: map[string]float64{},
	}
	return c
}

//format the cart
func (c *cart) format() string {
	fs := "Cart Breakdown \n"

	var total float64 = 0

	//list items
	for k, v := range c.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f","total:",total)
	return fs
}

//add an item to cart
func (c *cart) addItem(name string,price float64)  {
	c.items[name] = price
}

//save cart
func (c *cart) save(){
	data := []byte(c.format())

	err := os.WriteFile("cart/" +c.name+".txt",data,0644)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Cart details was saved to file")
}