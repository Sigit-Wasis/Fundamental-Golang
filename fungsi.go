package main

import"fmt"

func main() {
	/* fungsi adalah sekumpulan blog kode yang dibungkus dengan nama tertentu */
	// memanggil fungsi tampilkan pesan  
	tampilkan_pesan()

	// menampilkan pengembalian string
	fmt.Println(tampilkan_pesan())

	// memberi nilai parameter
	fmt.Println(tampilkan_pesan(10, 5))
}

// fungsi baru dengan nama tampilkan_pesan
func tampilkan_pesan(x int, y int)int {
	// mengambil nilai parameter
	var z = x + y
	return z

	// mengembalikan string
	return "Pesan telah Diterima"

	// Menampilan pesan 
	fmt.Println("Pesan Berhasil Diterima")
}