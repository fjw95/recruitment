package sort

import (
	"fmt"
)

// Function untuk menghasilkan tampilan visualisasi vertical barcharts
func Sorting(numArray []int) {

	high := 0
	// Looping untuk mengetahui nilai tertinggi dari array
	for i := 0; i < len(numArray); i++ {
		v := numArray[i]
		if v > high {
			high = v
		}
	}

	// Looping sebanyak nilai tertinggi, tiap selesai looping j dikurangi satu
	for j := high; j > 0; j-- {
		// Looping untuk menampilkan visualisasi sesuai dengan banyak nya nilai dalam array
		for k := 0; k < len(numArray); k++ {
			/*
				Apakah nilai index k == j
				Jika benar beri tanda "|"
				Jika tidak seleksi kondisi lagi
			*/
			if numArray[k] == j {
				fmt.Print(" | ")
			} else {
				/*
					Apakah nilai index k > j
					Jika benar beri tanda "|"
					Jika tidak tidak diberi tanda apa-apa
				*/
				if numArray[k] > j {
					fmt.Print(" | ")
				} else {
					fmt.Print("   ")
				}
			}
		}
		fmt.Println("")
	}
	// Looping untuk menampilkan nilai dari array
	for l := 0; l < len(numArray); l++ {
		// Tampilkan nilai dari index l
		fmt.Print(" ", numArray[l], " ")
	}
	fmt.Println("")
}
