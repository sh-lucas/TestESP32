package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/door-closed", Home)
	http.HandleFunc("/subscribe", Subscribe)

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	fmt.Println("Sever starting at :8080")
	http.ListenAndServe(":8080", nil)
}
