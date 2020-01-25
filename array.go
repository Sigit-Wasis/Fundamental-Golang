package main

import "fmt"

func main() {
	// Array (Kumpulan data bertipe data sama yang ada dalam suatu variabel)
	// array dengan nama variabel bauh dengan tipe data string dan jumlah data di dalam array 4
	var buah=[4]string{"apel", "jeruk", "anggur", "melon"}
	// untuk merubah index 1 (jeruk) menjadi (mangga)
	buah[1] = "mangga"
	// Mencetak isi dan elemet buah
	fmt.Println("Jumlah Element :", len(buah))
	fmt.Println("Isi Element :", buah)
}