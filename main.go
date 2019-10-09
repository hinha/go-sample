package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func tambah(x, y float64) float64 {
	return x + y
}

func multi_string(a, b string) (string, string) {
	return a, b
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Wow, running web")
}

// func memory_addressStr(a int) int {
// 	in_memory := &a

// 	return in_memory
// }

func main() {
	// Func tambah
	num1, num2 := 10.6, 6.8
	strA, strB := "Hello", "Welcome"

	var convertA int = 62
	var toconvert float64 = float64(convertA)

	names := []string{
		"Adams",
		"Peters",
		"Benght",
	}

	// Range Loop
	for i, name := range names {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	fmt.Println("Number rand 1-100", rand.Intn(100))
	fmt.Println(tambah(num1, num2))
	fmt.Println(multi_string(strA, strB))
	fmt.Println(toconvert)

	x := 100
	a := &x // memory address
	fmt.Println(a)
	fmt.Println(*a)

	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
