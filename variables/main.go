package main

import "fmt"

//somename := "bowser" No se pueden poner variables fuera de funciones

func main() {

	//string
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string
	nameThree = "peach"
	nameFour := "yoshi"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memoria

	//var numOne int8 = 215 esto no se puede porque int8 solo permite valores entre -128 y 128
	//var numOne int8 = 25
	//var numTwo  = - 128
	// var numThree unit = -128 esto no se puede por uint solo permite valores positivos

	//float

	var scoreOne float32 = -1.5
	var scoreTwo float64 = 888786876876876876
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
