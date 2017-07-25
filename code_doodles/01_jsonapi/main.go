package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	// UKey represents a unique identification key
	UKey = "oa4i§&2P1OJN%§5I%P"
)

type Dog struct {
	Name string
	Age  int
	Race string
}

type User struct {
	AuthToken string `json:"authtoken,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"fname,omitempty"`
	LastName  string `json:"lname,omitempty"`
	Password  string `json:"pwd,omitempty"`
}

type Account struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Number   string `json:"number,omitempty"`
}

func main() {
	port := flag.String("port", "8080", "an string")
	flag.Parse()

	http.HandleFunc("/", index)
	http.Handle("/api/dog", AuthWrapper(http.HandlerFunc(dog)))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// index tests simple api response
func index(w http.ResponseWriter, r *http.Request) {
	fruits := make(map[string]string)

	fruits["Apples"] = "2"
	fruits["Oranges"] = "1"
	fruits["YourIP"] = r.RemoteAddr
	fruits["AnotherIP"] = r.Header.Get("X-Forwarded-For")
	fruits["LastIP"] = r.Header.Get("X-Real-Ip")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fruits)
}

func dog(w http.ResponseWriter, r *http.Request) {
	d := Dog{Name: "Pluto", Age: 50, Race: "Celestial Dwarf"}
	json.NewEncoder(w).Encode(d)
}

func konto(w http.ResponseWriter, r *http.Request) {
	konto := Account{Name: "Cthulu", LastName: "OldOne", Number: "AT2093024802"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(konto)
}

func AuthWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing AuthWrapper logic before the Handler")
		state := CheckAuth(r)
		fmt.Printf("Authenticated: %t\n", state)
		if state {
			next.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}

		log.Println("Executing AuthWrapper logic after the Handler")
	})
}

func CheckAuth(r *http.Request) bool {
	d := json.NewDecoder(r.Body)
	var u User
	err := d.Decode(&u)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	// Get User From Database
	// Checking Authentication Token
	if u.AuthToken == UKey {
		return true
	}
	return false
}
