package main

import "fmt"

func main() {
	// Slice bersifat dinamis 
	// Array bersifat static
	var buah = []string{"apel", "mangga", "anggur", "jeruk"}
	// menambah elemen durian pada slice buah
	buah = append(buah, "durian")
	
	fmt.Println(buah)
	// len untuk menghitung jumlah elemen pada buah
	fmt.Println("Jumlah elemet", len(buah))
}