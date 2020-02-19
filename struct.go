package main

import"fmt"

func main() {
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), 
	// yang dibungkus sebagai tipe data baru dengan nama tertentu.
	var x1 pelajar
	var x2 guru

	x1.nama = "Sigit wasis subekti"
	x1.kelas = 1
	x1.umur = 19

	x2.nama = "Dwi Nur Hamid"
	x2.kelas = 1
	x2.umur = 19

	fmt.Println("nama :", x1.nama)
	fmt.Println("kelas :", x1.kelas)
	fmt.Println("umur :", x1.umur)

	fmt.Println("Gurunya adalah :")

	fmt.Println("nama :", x2.nama)
	fmt.Println("kelas :", x2.kelas)
	fmt.Println("umur :", x2.umur)
}

// Keyword type digunakan untuk deklarasi struct
type pelajar struct {
	nama string
	kelas int
	umur int
}

type guru struct {
	nama string
	kelas int
	umur int
}