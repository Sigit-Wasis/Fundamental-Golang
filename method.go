package main

import"fmt"

func main() {
	// Method adalah fungsi yang menempel pada struct
	var s1 pelajar
	var s2 = pelajar{"Hamid", 4}
	s1.nama = "Sigit Wasis"
	s1.umur = 19
	
	// Deklarasi nama method
	s1.namasaya()
	s2.namasaya()
}

// Membuat struct
type pelajar struct {
	nama string
	umur int
}

// Membuat Method
// s menyimpan semua data dari s1
func (s pelajar) namasaya() {
	fmt.Println("Nama saya adalah", s.nama)
	fmt.Println("Umur saya adalah", s.umur)
}