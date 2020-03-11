package main

import (
	"fmt"
	"strconv"
)

func greeting(name string) string {
	return "Hello " + name
}

//define a struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value receiver)

func (p Person) greet() string {
	fmt.Print(p)
	p.age++
	return "Hello my name is " + p.firstName + p.lastName + " And my age is " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver method)

func (p *Person) hasBirthday() {
	fmt.Print(p)
	// p.age++
}

func main() {

	person1 := Person{"Samanth", "sundar", "boston", "Male", 26}
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	fmt.Println(person1.age) //26 since value receiver wouldn't change the value at the address
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// fruit := [2]string{"Apple", "Bananna"}
	// veggieSlice := []string{"Brocoli", "Carrot", "Beans"}
	// var name = "Sai"
	// var num = 3
	// // i := 1
	// var isCool = true
	// const notCool = false
	// // notCool = true ---> can't be assigned as it's a const
	// name1, email := "Short notation  only inside a function", "test@mail.com"
	// fmt.Println("Hello, " + name + "!!!")
	// fmt.Println(name1, email)
	// fmt.Printf("%T\n", num)
	// fmt.Printf("%T\n", isCool)
	//
	// fmt.Println(fruit[0], fruit[1], veggieSlice)
	//
	// fmt.Println(math.Floor(3.676444))
	//
	// fmt.Println(greeting("go user"))
	// fmt.Println(strutil.Reverse("bye"))
	//
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	//
	// //FizzBuzz
	// for i := 1; i <= 100; i++ {
	// 	if i%15 == 0 {
	// 		fmt.Printf("Number is %d FizzBuzz\n", i)
	// 	} else if i%3 == 0 {
	// 		fmt.Printf("Number is %d Fizz\n", i)
	// 	} else if i%5 == 0 {
	// 		fmt.Printf("Number is %d Buzz\n", i)
	// 	} else {
	// 		fmt.Printf("Number is %d\n", i)
	// 	}
	// }
	//
	// //maps
	// emails := make(map[string]string)
	//
	// emails["test"] = "test@mail.com"
	// emails["testuser"] = "test@user.com"
	//
	// fmt.Println(emails["test"])
	//
	// ids := []int{23, 34, 4, 44, 24}
	// for i, id := range ids {
	// 	fmt.Printf("%d -ID : %d\n", i, id)
	// }
	//
	// // range returns two integers ID and the value assigned to it.
	// for _, id := range ids {
	// 	fmt.Printf("ID : %d\n", id)
	// }
	//
	// sum := 0
	// for _, id := range ids {
	// 	sum += id
	// }
	// fmt.Println("Sum ", sum)
	//
	// a := 5
	// b := &a
	//
	// fmt.Printf("%d and the address of where  it is stored %d\n", a, b)
	// fmt.Printf("type of the pointer %T\n", b)
	// fmt.Printf("Value at the address %d\n", *b)
	// *b = 35
	// fmt.Printf("%d did *b=35 which changed the value stored at  its address where a is pointing to and the address of where  it is stored %d\n", a, b)

}
