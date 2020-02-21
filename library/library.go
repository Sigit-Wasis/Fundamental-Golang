package library

import "fmt"

// Tulisan awal function kecil maka ini private
func iniprivate() {
	fmt.Println("saya di private")
}

// Tulisan awal function besar maka ini public
func Inipublic() {
	fmt.Println("saya di public")
	// jika ingin menjalankan yang private
	iniprivate()
}