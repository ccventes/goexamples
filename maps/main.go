package main

import (
	"fmt"
)

func Updatevalue(m map[string]float64, clave string, nuevoValor float64) {

	// Verificar si la clave existe en el mapa antes de modificarla
	if _, ok := m[clave]; ok { // ok es un valor booleano que es verdadero si el elemento se encuentra
		m[clave] = nuevoValor
	} else {
		fmt.Printf("La clave %s no existe en el mapa\n", clave)
	}
}

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	} // el map es como un diccionario, en este caso es una relacion de string -> float64 -> float64
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps

	for key, value := range menu { // inerando por cada key y tomando un valor "value" del menu

		fmt.Println(key, "-", value)

	}
	//update the menu
	for key, value := range menu {
		menu[key] = value + 5.0 // no imrpimir el cambio dentro del for
		fmt.Printf("Nuevo precio para %s: %.2f\n", key, menu[key])
		fmt.Println(key, "+", menu[key])

	}
	Updatevalue(menu, "pie", 3.48) //funcion para hacer update en el map
	for key, value := range menu { // inerando por cada key y tomando un valor "value" del menu

		fmt.Println(key, "(.)", value)

	}

}
