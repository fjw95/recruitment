package ans

import (
	"fmt"

	"github.com/galihrivanto/recruitment/sorting/sort"
)

// Function No. 1
func AnsOne(numArray []int) {
	// Menampilkan Visualisasi vertical barcharts
	sort.Sorting(numArray)
}

// Function No. 2
func AnsTwo(numArray []int) {
	// Menampilkan Visualisasi vertical barcharts array awal
	sort.Sorting(numArray)
	fmt.Println("Step 0 (Tampilan array awal)")
	count := 0
	// Looping sebanyak nilai yang ada pada array
	for i := 0; i < len(numArray); i++ {
		// Looping utk melakukan insertion sort
		for j := i; j > 0 && numArray[j] < numArray[j-1]; j-- {
			count++
			// Pindah index
			tmp := numArray[j]
			numArray[j] = numArray[j-1]
			numArray[j-1] = tmp
			// Menampilkan Visualisasi vertical barcharts
			sort.Sorting(numArray)
			fmt.Println("Step", count, "(Pindah posisi antara", numArray[j], "dan", numArray[j-1], ")\n")
		}
	}
}

// Function No. 3
func AnsThree(numArray []int) {
	// Menampilkan Visualisasi vertical barcharts array awal
	sort.Sorting(numArray)
	fmt.Println("Step 0 (Tampilan array awal)")
	count := 0
	// Looping sebanyak nilai yang ada pada array
	for i := 0; i < len(numArray); i++ {
		// Looping utk melakukan insertion sort secara descending
		for j := i; j > 0 && numArray[j] > numArray[j-1]; j-- {
			fmt.Println(numArray[j], numArray[j-1])
			count++
			tmp := numArray[j]
			numArray[j] = numArray[j-1]
			numArray[j-1] = tmp
			// Menampilkan Visualisasi vertical barcharts
			sort.Sorting(numArray)
			fmt.Println("Step", count, "(Pindah posisi antara", numArray[j], "dan", numArray[j-1], ")\n")
		}
	}
}
