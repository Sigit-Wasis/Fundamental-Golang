package main

import "fmt"

func main() {
	// Contoh Perulangan Pertama
	for i := 0; i <= 10; i++ {
		fmt.Println("Angka", i)
	}

	// Contoh Perulangan Kedua
	var i = 0
	for i <= 10 {
		fmt.Println("Angka", i)
		i++
	}

	// Contoh Perulangan Ketiga
	var i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 10 {
			break;
		}
	}	
}