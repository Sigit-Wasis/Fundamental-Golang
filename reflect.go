package main

import "fmt"
import "reflect"

func main() {
	var number = "sigit"
	var reflectnumber = reflect.ValueOf(number)

	// melihat type data dari variabel number
	fmt.Println("Tipe Variabel :", reflectnumber.Type())

	// Kind mengecek dahulu tipe data 
	// Jika integer maka tampikan nilai variabelnya
	if reflectnumber.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel :", reflectnumber.Int())
	// Jika string maka tampikan nilai variabelnya
	} else if reflectnumber.Kind() == reflect.String {
		fmt.Println("Nilai Variabel :", reflectnumber.String())
	}
}