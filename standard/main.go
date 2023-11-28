package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greetings := "hello there friends"
	fmt.Println(strings.Contains(greetings, "hello"))
	//fmt.Println(strings.ReplaceAll(greetings, "hello", "Hi"))
	//fmt.Println(strings.ToUpper(greetings)) Mayusculas
	//fmt.Println(strings.Index(greetings, "hello"))//busca el indice dondecomienza la incidencia
	//fmt.Println(strings.Split(greetings, " "))

	//sort organizar

	ages := []int{7, 2, 3, 4, 5, 21, 7, 8, 9, 55, 11, 12, 6, 14, 8}
	index := sort.SearchInts(ages, 7) //encuentra el indice del primer elemento encontrado dado por el usuario en una listaq ordenada
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(index)
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))

}
