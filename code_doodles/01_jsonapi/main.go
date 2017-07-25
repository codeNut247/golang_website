package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	jsonloader "github.com/codeNut247/golang_website/code_doodles/01_jsonapi/app/jsonloader.go"
)

const (
	// UKey represents a unique identification key
	UKey = "dripler"
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

<<<<<<< HEAD
type Account struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Number   string `json:"number,omitempty"`
=======
type LocalContext struct {
	userSlice []User
}

var LocalDB LocalContext

func (lc *LocalContext) ParseJson(b []byte) error {
	return json.Unmarshal(b, &lc)
}

func init() {
	jsonloader.Load("users.json", LocalDB)
>>>>>>> f3b359a11d2fb078fcd54998d5b9de627c97e767
}

func main() {
	fmt.Println(LocalDB)
	port := flag.String("port", "8080", "a string")
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
<<<<<<< HEAD
		state := CheckAuth(r)
		fmt.Printf("Authenticated: %t\n", state)
		if state {
			next.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}

=======

		if CheckAuth(r) {
			log.Println("Authorized")
		} else {
			log.Println("Not Authorized")
		}

		next.ServeHTTP(w, r)

>>>>>>> f3b359a11d2fb078fcd54998d5b9de627c97e767
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
