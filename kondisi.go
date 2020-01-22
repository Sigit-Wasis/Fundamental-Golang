package main

import "fmt"

func main() {

	// Pengkondisian If - else
	var point = 8

	if point == 10 {
		fmt.Println("Lulus dengan Sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	// Pengkondisian Switch - case
	// Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel, lalu kemudian di-cek nilainya.
	var exam = 10

	switch exam {
		case 8:
			fmt.Println("Perfect")
		case 7:
			fmt.Println("Kurang Baik")
		case 6:
			fmt.Println("Kurang")
		default:
			fmt.Println("No bad")
	}

	// Penggunaan switch dengan case untuk banyak kondisi
	var tugas = 100

	switch tugas {
		case 90, 80, 100:
			fmt.Println("Good")
		case  40:
			fmt.Println("Jelek")
		default:
			fmt.Println("Harus Belajar Lagi")
	}

	// Seleksi Kondisi Bersarang
	var latihan = -1

	if latihan > 7 {
		switch latihan {
			case 10:
				fmt.Println("Perfect")
			default:
				fmt.Println("Nice")
		}
	} else {
		if latihan == 5 {
			fmt.Println("Not bad")
		} else if point == 3 {
			fmt.Println("Tidak Belajar")
		} else {
			fmt.Println("Gk pernah buka Buku")

			if latihan == 0 {
				fmt.Println("Coba Lagi Bro")
			} else if latihan == -1{
				fmt.Println("Keterlaluan")
			}
		}
	}
}