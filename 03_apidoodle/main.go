package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	UKey = "oaij%JOAO%PO§O$pGQO§J%GOGJQPOJN%§I%P"
)

type Dog struct {
	Name string
	Age  int
	Race string
}

type User struct {
	AuthKey   string
	FirstName string
	LastName  string
	Password  string
}

func main() {
	http.HandleFunc("/", index)
	handler := http.HandlerFunc(dog)
	http.Handle("/api/dog", AuthWrapper(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// index tests simple api response
func index(w http.ResponseWriter, r *http.Request) {
	fruits := make(map[string]int)

	fruits["Apples"] = 2
	fruits["Oranges"] = 10

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fruits)
}

func dog(w http.ResponseWriter, r *http.Request) {

}

func AuthWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing AuthWrapper logic before the Handler")
		next.ServeHTTP(w, r)
		log.Println("Executing AuthWrapper logic after the Handler")
	})
}
