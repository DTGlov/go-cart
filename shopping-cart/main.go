package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func getInput(prompt string,r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input,err := r.ReadString('\n')

	return strings.TrimSpace(input),err
}

func createCart() cart {
	reader := bufio.NewReader(os.Stdin)

	name,_ := getInput("Create a new cart name: ",reader)

	c:= newCart(name)
	fmt.Println("Created cart for -", c.name)

	return c

}


func promptOptions(c cart){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := getInput("Choose option (a - add item, s - save cart): ",reader)
	
	switch opt{
	case "a" :
		name,_ := getInput("Item name: ",reader)
		price,_ := getInput("Item price: ",reader)

		p,err := strconv.ParseFloat(price,64)
		if err !=nil {
			fmt.Println("The price must be a number")
			promptOptions(c)
		}
		c.addItem(name,p)
		fmt.Println("item added - ",name,price)
		promptOptions(c)
	case "s":
		c.save()
		fmt.Println("you saved the file - ",c.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(c)
	}

}


func main() {

	mycart := createCart()
	promptOptions(mycart)
}