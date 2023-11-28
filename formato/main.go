package main

import "fmt"

func main() {
	age := 35
	name := "shaun"

	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n ")

	fmt.Println("Hello ninjas!")
	fmt.Println("Goodbye ninjas!")
	fmt.Println("my age is ", age, "and my name is ", name)

	//printf print con formato
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name) //quote " "
	fmt.Printf("age is of type %T \n", age)                    //tipo
	fmt.Printf("you scored %0.1f points \n", 255.55)           //float
	//Sprintf para salvar strings con formato
	str := fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Printf(str)

}
