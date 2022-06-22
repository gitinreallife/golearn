package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("Hi bitch, welcome to Go 102")
	// //===function as value
	// functionValue := getFunctionValue
	// fmt.Println(functionValue("Bad person"))

	// sayHelloWithFilter("gita", spamFilter)
	// filter := spamFilter
	// sayHelloWithFilter("bitch", filter)

	// sayHelloWithFilterType("type", spamFilter)

	// //anonymous function
	// nameLuar := "bukan anonymous"
	// blacklist := func(name string) bool {
	// 	nameLuar = "changed"
	// 	return name == "bitch"
	// }
	// registerUser("admin", blacklist)
	// registerUser("bitch", func(name string) bool {
	// 	return name == "bitch"
	// })
	// //check variable that is changed in anonymous function
	// fmt.Println("----", nameLuar)

	fmt.Println("factorial recursive-------")
	recursive := factorialRecursive(4)
	fmt.Println(recursive)
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked. ", name)
	} else {
		fmt.Println("Welcome, ", name)
	}
}

// function type declaration
type Filter func(string) string

func sayHelloWithFilterType(name string, filter Filter) {
	fmt.Println("HELLO TYPE ", filter(name))
}
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("HELLO ", filter(name))
}
func spamFilter(name string) string {
	if name == "bitch" {
		return "..."
	} else {
		return name
	}
}
func getFunctionValue(name string) string {
	fmt.Println("")
	fmt.Println("---Function as Value---")
	return "bye " + name
}
