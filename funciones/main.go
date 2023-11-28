package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		println("v: ", v)

		initials = append(initials, v[:1]) // con los strings no se puede  acceder a un solo eleemento, sino con slices a:b
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {

	i, j := getInitials("tifa lockheart")

	fmt.Println("las iniciales son", i, j)
	aleatorio := sayRandomName()
	println("nombre aleatorio ", aleatorio)
	k, l := getInitials(aleatorio)
	fmt.Println("las iniciales son", k, l)

}
