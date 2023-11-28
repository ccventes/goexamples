package main

// asociación de funciones con structuras para hacer POO
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n') //espera un newline (ENTER)

	return strings.TrimSpace(input), error

} // recibe como parametro el puntero del reader que estamos usando para no crear una copia
func createbill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b

}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
	case "a":
		name, _ := getInput("Iten name ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)

		}
		b.addItem(name, p)
		fmt.Println("item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)

		}

		b.UpdateTip(t)
		promptOptions(b)
		fmt.Println("tip added: ", tip)
	case "s":
		b.save()
		fmt.Println("you chose to save the bill", b)

	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)

	}

}

func main() {

	mybill := newBill("Mario's bill")
	mybill.addItem("onion soup", 4.20)
	mybill.addItem("Veg Pipe", 5.20)
	mybill.addItem("Tofee Pudding", 4.30)
	mybill.addItem("Cofee", 4.10)
	mybill.UpdateTip(10)

	fmt.Println(mybill.format())

	//Otro modo de crear biils
	myOtherBill := createbill()
	fmt.Println(myOtherBill)
	promptOptions(myOtherBill)

}
