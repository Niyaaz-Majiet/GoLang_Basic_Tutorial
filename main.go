package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Hello World 
	fmt.Println("Hello from a module, Ledgends!");


    // Declare variable and print
	var i int = 42;

	fmt.Println("Value : "+ strconv.Itoa(i));

	var x float32 = 42.32;

	fmt.Println("Value : "+ fmt.Sprintf("%f", x));

	faveFood := "Chicken Curry";

	fmt.Println("Favourite Food :" + faveFood);

	c:= complex(3,4);

	fmt.Println(c);

	//Working with pointers
	//Variable is called a pointer type because or variable will point to another location that will hold our data 


	var firstName *string = new(string);
	*firstName = "Niyaaz - Majiet"

	fmt.Println("Name : " + *firstName);
    

}