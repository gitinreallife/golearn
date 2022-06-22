package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	fmt.Println("Hi bitch, welcome to Go 101")
	//basic bich
	// // tipeData()
	// // tipeVariable()
	// // conversion()
	// mathOperation()
	//basic bich

	// tipeDataDeclarationAliases()
	// tipeDataMap()
	tipeDataArray()
	// tipeDataSlice()
	// condition()
	// counterLoop()

	// sums(1, 2, 3)
	// sums([]int{1, 1, 1, 1, 1}...)
	// nums := []int{9, 10, 10, 20}
	// sums(nums...)

	fmt.Println("")
	fmt.Println("######ENDS HERE#######")

}
func sums(nums ...int) {
	fmt.Println("")
	fmt.Println("---Variadic Function---")
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func multipleReturn(param1 string) (firstReturn string, secondOne bool) {
	return "bithc", true
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
func condition() {
	fmt.Println("")
	fmt.Println("---Condition basic bich---")
	name := "gita"
	if name == "bruh" {
		fmt.Println(name)
	} else if name == "gita" {
		fmt.Println("eheheheh")
	} else if name == "gita" {
		fmt.Println("eheheheh 2")
	} else {
		fmt.Println("tf")
	}
	fmt.Println("")
	fmt.Println("---Condition if short statement---")

	if length := len(name); length < 5 {
		fmt.Println("name is too short: ", length)
	} else {
		fmt.Println("just fine: ", length)
	}

	switch name {
	case "bruh":
		fmt.Println(name)
	default:
		fmt.Println("default bich")
		fmt.Println("default bich 2")
	}

	//short switch
	switch lengths := len(name); lengths > 5 {
	case true:
		fmt.Println(name)
	default:
		fmt.Println("name is too short: ", lengths)
	}

	fmt.Println("no condition switch")
	switchLength := len(name)
	switch {
	case switchLength > 10:
		fmt.Println("too long bitch")
	case switchLength < 5:
		fmt.Println("pas")
	case switchLength < 3:
		fmt.Println("cek deh")
	default:
		fmt.Println("wtf")

	}
	// fmt.Println("print length: ", length)
}

func tipeDataDeclarationAliases() {
	//check declaration aliases builtin.go
	// byte()
	fmt.Println("")
	fmt.Println("---Tipe Data Declaration Aliases---")
	type NoKTP string
	var noKtp NoKTP = "1234567890123456"

	type IDCard struct {
		noKtp       NoKTP
		age         int
		city, phone string
	}

	gita := IDCard{
		noKtp: noKtp,
		age:   99,
		city:  "Kochi",
		phone: "+62123456789",
	}

	fmt.Println("noKtp:", gita.noKtp)

	id2 := IDCard{noKtp, 10, "bruh", "+62165412333"}
	fmt.Println(`inline declaration id2 := IDCard{noKtp, 10, "bruh", "+62165412333"}: NoKTP: `, id2.noKtp)
}

func tipeDataSlice() {
	fmt.Println("")
	fmt.Println("---Tipe Data Slice---")
	fmt.Println(`Slice memiliki 3 data, yaitu pointer, length and capacity.
	pointer adalah penunjut data pertama di array para slice
	length adalah panjang dari slice
	capacity adalah kapasitas dari slice dimana length <= capacity
	`)
	months := [...]string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sept", "oct", "nov", "dec"}
	monthSliceSummer := months[5:9]
	monthSliceSummer[0] = "changedinSlice"
	fmt.Println("Does it change the array: ", months)
	// pointer = 5
	// length = 4
	// capacity = 7
	fmt.Println("monthSliceSummer: ", monthSliceSummer)
	fmt.Println("monthSliceSummer lenght: ", len(monthSliceSummer))
	fmt.Println("monthSliceSummer capacity: ", cap(monthSliceSummer))

	monthSliceSummer = append(monthSliceSummer, "newappendedmonth")
	monthSliceSummer = append(monthSliceSummer, "newappendedmonth2")
	// fmt.Println("monthSliceNewSummer: ", monthSliceNewSummer)
	fmt.Println("months: ", months)
	fmt.Println("monthSliceSummer: ", monthSliceSummer)

}
func tipeDataArray() {
	fmt.Println("")
	fmt.Println("---Tipe Data Array---")
	names := [3]string{"gita", "arifatun", "nisa"}
	names[1] = "changed"
	fmt.Println(`inisiasisasi names := [3]string{"gita", "arifatun", "nisa"}
	names[1]: `, names[1])

	unlimitedArray := [...]string{"XS", "S", "M", "L"}
	fmt.Println(`inisiasisasi unlimitedArray := [...]string{"XS", "S", "M", "L"}
	unlimitedArray[1]: `, unlimitedArray[1])
	fmt.Println("len(unlimitedArray): ", len(unlimitedArray))
}
func mathOperation() {
	fmt.Println("")
	fmt.Println("---Math Operation---")
	a := 100
	a += 20
	a -= 10
	a *= 2
	a /= 5
	a %= 5
	fmt.Println("hasilnya: ", a)
	fmt.Println("")
	fmt.Println("---Unary Operator---")
	a++
	a--
	fmt.Println("hasilnya: ", a)
}
func tipeData() {
	fmt.Println("")
	fmt.Println("---Tipe Data Number---")
	var integer8 int8 = -128
	fmt.Println("integer terkecil: ", integer8)
	var unsigned8 uint8 = 255
	fmt.Println("unsigned integer 8 terbesar: ", unsigned8)
	var floating32 float32 = 3.5
	fmt.Println("floating 8: ", floating32)

	var aliasbyte byte = 255
	fmt.Println("alias byte = uint8 = ", aliasbyte)
	var aliasrune rune = 32
	fmt.Println("alias rune = int32 = ", aliasrune)
	var aliasint int = 32
	fmt.Println("alias int based on OS, min int32 = ", aliasint)

	fmt.Println("")
	fmt.Println("---------Tipe Data Boolean---------")
	var booltrue bool = true
	fmt.Println("bool true: ", booltrue)
	var boolfalse bool = false
	fmt.Println("bool false: ", boolfalse)

	fmt.Println("")
	fmt.Println("---------Tipe Data string---------")
	var stringgo string = "ab£"
	fmt.Println("stringgo bich: ", stringgo)
	fmt.Println("output string symbolnya len(stringgo) : ", len(stringgo))
	lengthh := len([]rune(stringgo))
	fmt.Println("real length len([]rune(stringgo)): ", lengthh)
	fmt.Println("")

	fmt.Println("")
	var multilinestring string = "this string is on \n" +
		"multiple line bitch \n" +
		"ehe kan bener."
	fmt.Println(multilinestring)
	fmt.Println("")
	var rawstringliterals string = `escaping all " ' \" `
	fmt.Println("Raw String literals using `: ", rawstringliterals)
}

func tipeVariable() {

	fmt.Println("")
	fmt.Println("--------- Variable ---------")
	fmt.Println(` var bruh string = "stringvalue"`)
	fmt.Println(` bruh := "string value `)
	var (
		firstvarstring  = "string first"
		secondintstring = 32
	)
	fmt.Println(` 
	var (
		firstvarstring = "string first"
		secondintstring = 32
	)
	`)
	fmt.Println("testprint first varstring: ", firstvarstring)
	fmt.Println("testprint secondintstring: ", secondintstring)

	fmt.Println("")
	fmt.Println("--------- Constant ---------")
	const firstcontant string = "eko"
	const (
		secondconstant = "secondconstant"
		thirdconstant  = 3
	)
	fmt.Println(`const (
		secondconstant = "secondconstant"
		thirdconstant  = 3
	)`)
}

func conversion() {
	fmt.Println("")
	fmt.Println("--------- Konversi  ---------")
	var nilai32 int32 = 1000000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai32: ", nilai32)
	fmt.Println("int64(nilai32): ", nilai64)
	fmt.Println("hasil overflow int8(nilai32): ", nilai8)
}
