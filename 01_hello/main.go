package main

import "fmt"

func main() {
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
}
