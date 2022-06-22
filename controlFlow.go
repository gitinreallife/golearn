package main

import "fmt"

func logging(message string, controlType string) {
	fmt.Println("Logging function.", controlType+": "+message)
}
func main() {
	// deferFunction()
	// panicFunction(true)

	defer recoverFunction()
	panicFunction(true)
}

//recover: menangkap data panic sehingga program akan tetap berjalan
func recoverFunction() {
	message := recover()
	fmt.Println("Error? ", message)
	logging("1", "recover")
}

func panicFunction(error bool) {
	if error {
		panic("App error")
	}
	fmt.Println("App is running")
}

func deferFunction() {
	//first in last out
	defer logging("(1)", "defer")
	defer logging("(2)", "defer")
	defer logging("(3)", "defer")

	fmt.Println("run app")
	//throw error
	divint := 0
	result := 10 / divint
	fmt.Println(result)
}
