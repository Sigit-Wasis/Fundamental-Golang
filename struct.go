package main

import"fmt"

func main() {
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), 
	// yang dibungkus sebagai tipe data baru dengan nama tertentu.
	var x1 pelajar
	x1.nama = "Sigit wasis subekti"
	x1.kelas = 1
	x1.umur = 19

	fmt.Println("nama :", x1.nama)
	fmt.Println("kelas :", x1.kelas)
	fmt.Println("umur :", x1.umur)
}

type pelajar struct {
	nama string
	kelas int
	umur int
}