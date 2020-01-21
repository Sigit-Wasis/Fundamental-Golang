package main

import "fmt"

func main() {
	// Konstanta (Nilai yang tidak dapat diubah)
	const x = 30

	const firstname string = "Sigit Wasis Subekti"

	// Perbedaan Println dengan Print
	/* 
		Println = menghasilkan baris baru di akhir outputnya & nilai parameter digabung dengan penghubung spasi
		Print = tidak menghasilkan barus baru di akhir outputnya & nilai parameter di gabung tanpa pemisah
	*/
	fmt.Println(x)
	fmt.Print(firstname)
}