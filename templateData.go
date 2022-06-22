package main

import "fmt"

func main() {
	// tipeDataMap(false)
	tipeDataStruct()
}

/**
Map key dan value nya harus seragam
struct tidak perlu seragam
*/
//semacam class dan attributes in golang

type NoKTP string
type IDCard struct {
	noKtp       NoKTP
	age         int
	city, phone string
}

//add method ke struct, struct method
func (ktp IDCard) printDeets() {
	fmt.Println(ktp)
}

func tipeDataStruct() {
	var noKtp NoKTP = "1234567890123456"
	//best practie declaration struct
	gita := IDCard{
		noKtp: noKtp,
		age:   99,
		city:  "Kochi",
		phone: "+62123456789",
	}
	gita.city = "changed"
	gita.printDeets()
}

func tipeDataMap(isGetMap bool) (map[int]string, bool, string) {
	if !isGetMap {
		fmt.Println("")
		fmt.Println("---Tipe Data Map---")
		person := map[string]string{
			"name": "gita",
			"age":  "99",
		}
		person["title"] = "proffessional sloucher"
		fmt.Println(person)
		delete(person, "title")
		fmt.Println(person)
		fmt.Println(person["age"])
	}

	ranking := map[int]string{
		1: "gita",
		2: "arifatun",
		3: "nisa",
	}
	fmt.Println(ranking)
	return ranking, true, "bitch"
}
func counterLoop() {
	counter := 1
	for i := 0; i < counter; i++ {
		fmt.Println(i)
	}
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
	slice := []string{"gita", "arifatun", "nisa"}
	for i := 0; i < len(slice); i++ {
		fmt.Println((slice[i]))
	}

	fmt.Println("foreach ---------")
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
	for _, value := range slice {
		fmt.Println("declare _ for unused index: ", value)
	}

	fmt.Println("foreach map---------")
	rankingMap, bitch, _ := tipeDataMap(true)
	// keys := make([]string, 0, len(m))
	// sort.Ints(rankingMap)
	//how to sort
	fmt.Println("get second values tipedatamap ", bitch)
	for key, value := range rankingMap {
		if key == 1 {
			fmt.Println("break at ranking ", key, " si ", value)
			break
		} else if key == 2 {
			//skip setelahnya
			continue
		} else {
			fmt.Println("ranking ", key, " adalah ", value)
		}
		fmt.Println("ranking check continue")

	}
}
