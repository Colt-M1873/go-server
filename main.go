package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf(("Starting go miniserver at port 8080\n"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "handling url error: 404 not fount", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "handling woring methods method not supported except for GET", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello from go mini server!")
	fmt.Println("helloHandler() is called")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Printf("formHandler() is called: \n Name = %s\n Address = %s\n", name, address)
}
