package main

import (
	"fmt"

	"github.com/galihrivanto/recruitment/sorting/ans"
)

func main() {
	// Local variabel array
	var numArray = []int{1, 4, 5, 6, 8, 2}

	// Hasil No. 1
	fmt.Println("1. Visualisasi")
	// Menjalankan program No. 1
	ans.AnsOne(numArray)
	fmt.Println()

	// Hasil No. 2
	fmt.Println("2. Visualisasi dan Implementasi Insertion Sort")
	// Menjalankan program No. 2
	ans.AnsTwo(numArray)
	fmt.Println()

	numArray = []int{1, 4, 5, 6, 8, 2}
	// Hasil No. 3
	fmt.Println("3. Visualisasi dan Implementasi Reverse Insertion Sort (Descending)")
	// Menjalankan program No. 3
	ans.AnsThree(numArray)
}
