package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
params
-- How to declare variables
**/

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add Item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name ", reader)
		price, _ := getInput("Item Price ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOption(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOption(b)
	case "t":
		tip, _ := getInput("Enter tip amount ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOption(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOption(b)
	case "s":
		b.save()
		fmt.Println("You saved the file - ", b.name)
	default:
		fmt.Println("Your answer not a valid option...")
		promptOption(b)
	}
}

func main() {
	// Var declaration
	// var name string = "Christian Chiemela"
	// name = "Christian Chiemela"
	// var names string
	// // Don't declare shorthand outside of the func it won't work
	// named := "Christian Chiemela"
	// fmt.Println(name, names, named)

	// // integer
	// var numberOne int = 2
	// numberTwo := 4
	// fmt.Println(numberOne, numberTwo)

	// // bit & memory
	// var numberThree int8 = 127
	// // unit not negative number
	// var numberFour uint = 25
	// fmt.Println(numberThree, numberFour)

	// // Arrays

	// var nameone [3]string = [3]string{"hello", "body", "new"}

	// var nametwo = []int{1, 34, 5}
	// namethree := append(nametwo, 50)
	// fmt.Println(nameone, nametwo, namethree)

	// bill := newBill("christian chiemela")

	// bill.addItem("cake", 5.50)
	// bill.addItem("tea", 3.50)
	// bill.addItem("bread", 2.50)

	// bill.updateTip(10)

	// fmt.Println(bill.format())

	myBill := createBill()

	promptOption(myBill)

	// fmt.Println(myBill)

}
