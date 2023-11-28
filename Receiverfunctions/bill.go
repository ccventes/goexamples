package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b

}

// Parentesis a la izquierda son usados para declarar el receptor
func (b bill) format() string {

	fs := "Bill Breakdown: \n"
	var total float64 = 0.0
	// list items in bill
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...$%v \n", k+":", v)
		total += v
	}

	//total %-25v pone el valor 25 posiciones a la derecha
	fs += fmt.Sprintf("%-25v...$%0.2f", "total:", total+b.tip)
	return fs

}

// Update tip
func (b *bill) UpdateTip(tip float64) {
	b.tip = tip

}

// agregar un item a la cuenta
func (b *bill) addItem(name string, price float64) {

	b.items[name] = price
}

// Manejo de archivos

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile(b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")

}
