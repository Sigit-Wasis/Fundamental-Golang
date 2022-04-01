package main

import "fmt"

func main() {
	// Variabel dengan menuliskan tipe data
	var firstname string = "Sigit"
	// Variabel tanpa menuliskan tipe data
	var lastname = "Wasis Subekti"
	// Variabel tanpa menuliskan var
	address := "Lampung"
	// Mengubah varibel address
	address = "Lambar"
	// Digunakan untuk menampilan test
    fmt.Println("Nama Saya :", firstname, lastname, "Alamat :", address)

	/**
	 - var di atas digunakan untuk deklarasi variabel.
	 - skemanya adalah var <nama variabel> <tipe-data>
	 - fmt.Printf digunakan untuk untuk menampilkan output dalam bentuk tertentu
	 - fmt.Println digunakan untuk menampilkan output yang secara otomatis menghasilkan new line (baris baru) di akhir.
	 - tanda + untuk penggabungan string
	 - \n, untuk memunculkan baris baru di akhir
	*/

	// ---- DEKLARASI DENGAN TIPE DATA ----
	var firstName string = "Sigit"

	var lastName string
	lastName = "Wasis Subekti"

	fmt.Printf("halo john wick!\n")
	fmt.Printf("hello %s %s \n", firstName, lastName);
	fmt.Println("halo", firstName, lastName + "!")

	/**
	- lastNameBrother dideklarasikan dengan menggunakan metode type inference.
	- Tipe data lastNameBrother secara otomatis akan ditentukan menyesuaikan value atau nilai-nya.
	- Tanda := hanya digunakan sekali di awal pada saat deklarasi.
	- Perlu diketahui, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.
	*/


	// --- DEKLARASI TANPA TIPE DATA ---
	var firstNameBrother string = "dwi"
	lastNameBrother := "nur hamid"

	fmt.Printf("Hai %s %s!\n", firstNameBrother, lastNameBrother);


	// --- DEKLARASI MULTIPLE VARIABEL ---
	// cara pertama
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// cara kedua
	var fourth, fiveth, sixth = "empat", "lima", "enam"

	// cara ketiga
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Println(first, second, third)
	fmt.Println(fourth, fiveth, sixth)
	fmt.Println(seventh, eight, ninth)

}