package main

import (
	"fmt"
	"strconv"
	"github.com/niyaaz-majiet/webservice/models"
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

	//const

	const pi = 3.1415

	fmt.Println("Constant Var : "+fmt.Sprintf("%f", x));

	//Array
	var arr [3]int;
	arr[0] = 1;
	arr[1] = 2;
	arr[2] = 3;

    s := fmt.Sprint(arr);
    fmt.Printf("%q", s);
    fmt.Println("");

    // Slice data type
    slice:= []int{1,2,3};

    fmt.Println(slice);

       //adding new data to slice
    slice = append(slice,4,42,27)
    fmt.Println(slice);

      //slices of slices
    slice2:= slice[1:]
    slice3:= slice[:2]
    slice4:= slice[1:2]

    fmt.Println(slice2,slice3,slice4);


   //maps

   m := map[string]int{"foo":42}
   fmt.Println(m);
   fmt.Println(m["foo"]);

   //maps - update
   m["foo"] = 28;
   fmt.Println(m)

   //maps - delete
   delete(m,"foo");
   fmt.Println(m)

   //struct
   type user  struct {
     ID int
     FirstName string
     LastName string
   }

   var user1 user
   user1.ID = 1;
   user1.FirstName = "Niyaaz"
   user1.LastName = "Majiet"
   fmt.Println(user1);

   user2 := user{ID:2,FirstName:"Sam",LastName:"Pie"}
   fmt.Println(user2);

   //Struts imported

   u:= models.User{
      ID: 3,
      FirstName: "Max",
      LastName: "King",
   }

   fmt.Println(u);


}


