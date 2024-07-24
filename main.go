package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
params
-- How to declare variables
**/

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
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

}
