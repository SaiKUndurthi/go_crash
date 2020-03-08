package main

import (
	"fmt"
	"math"

	"github.com/SaiKUndurthi/go_crash/01_hello/strutil"
)

func greeting(name string) string {
	return "Hello " + name
}

func main() {

	fruit := [2]string{"Apple", "Bananna"}
	veggieSlice := []string{"Brocoli", "Carrot", "Beans"}
	var name = "Sai"
	var num = 3
	var isCool = true
	const notCool = false
	// notCool = true ---> can't be assigned as it's a const
	name1, email := "Short notation  only inside a function", "test@mail.com"
	fmt.Println("Hello, " + name + "!!!")
	fmt.Println(name1, email)
	fmt.Printf("%T\n", num)
	fmt.Printf("%T\n", isCool)

	fmt.Println(fruit[0], fruit[1], veggieSlice)

	fmt.Println(math.Floor(3.676444))

	fmt.Println(greeting("go user"))
	fmt.Println(strutil.Reverse("bye"))

}
