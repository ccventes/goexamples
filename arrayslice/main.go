package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 25, 35} // declaración de un arreglo, siempre tiene que ser de tamaño fijo
	var ages = [3]int{20, 25, 35}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	names[1] = "luigi"

	//slices arreglos que permiten agregar elementos
	var scores = []int{100, 50, 60}
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))
	fmt.Println(names, len(names))

	//slice ranges (porciones del slice)
	rangeOne := names[1:3]
	fmt.Println(rangeOne, len(rangeOne))

}
