package main

import"fmt"

func main() {
	// map adalah tipe data asosatif pada go berbentuk key-value pair
	var harga_makanan = map[string]int{"ayam_goreng":20000,"sate_ayan":15000}
	fmt.Println("ayam goreng", harga_makanan["ayam_goreng"])

}