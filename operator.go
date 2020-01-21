package main

import "fmt"

func main() {

	// Operator Aritmatika
	/*
		Penjumlahan (+)
		Pengurangan (-)
		Perkalian (*)
		Pembagian (/)
		Modulus (%)
	*/
	var kelasA = 10
	var kelasB = 5

	fmt.Println("Penjumlahan kelasA + kelasB :", kelasA + kelasB)
	fmt.Println("Pengurangan kelasA - kelasB :", kelasA - kelasB)
	fmt.Println("Perkalian kelasA * kelasB :", kelasA * kelasB)
	fmt.Println("Pembagian kelasA / kelasB :", kelasA / kelasB)
	fmt.Println("Modulus kelasA & kelasB :", kelasA % kelasB)

	// Operator Perbandingan 
	/*
		Sama dengan (==)
		Tidak sama dengan (!=)
		Lebih kecil daripada (<)
		Lebih kecil atau sama dengan (<=)
		Lebih Besar daripada (>)
		Lebih Besar atau sama dengan (>=)
	*/
	var kelasC = 60 / 2 + 5
	var hasil = kelasC ==  35

	fmt.Println("Nilai :", hasil)

	// Operator Logika
	/*
		kiri Dan kanan (&&)
		kiri Atau kanan (||)
		kebalikan (!)
	*/
	var bendaA = false
	var bendaB = true
	var bendaC = !true

	fmt.Println(bendaA && bendaB)
	fmt.Println(bendaA || bendaB)
	fmt.Println(bendaC)
}