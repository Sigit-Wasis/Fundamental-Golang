package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "apa kabar!")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "halo!")
    })

    // Routing aplikasi web (penentuan aksi ketika url di akses)
    http.HandleFunc("/index", index)

    fmt.Println("starting web server at http://localhost:8080/")
    // Digunakan untuk menghidupkan dan menjalankan go
    http.ListenAndServe(":8080", nil)
}