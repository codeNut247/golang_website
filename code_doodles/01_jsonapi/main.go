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

type LocalContext struct {
	userSlice []User
}

var LocalDB LocalContext

func (lc *LocalContext) ParseJson(b []byte) error {
	return json.Unmarshal(b, &lc)
}

func init() {
	jsonloader.Load("users.json", LocalDB)
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
	fruits := make(map[string]int)

	fruits["Apples"] = 2
	fruits["Oranges"] = 10

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(fruits)
}

func dog(w http.ResponseWriter, r *http.Request) {
	d := Dog{Name: "Pluto", Age: 50, Race: "Celestial Dwarf"}
	json.NewEncoder(w).Encode(d)
}

func AuthWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing AuthWrapper logic before the Handler")

		if CheckAuth(r) {
			log.Println("Authorized")
		} else {
			log.Println("Not Authorized")
		}

		next.ServeHTTP(w, r)

		log.Println("Executing AuthWrapper logic after the Handler")
	})
}

func CheckAuth(r *http.Request) bool {
	d := json.NewDecoder(r.Body)
	var u User
	err := d.Decode(&u)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	// Get User From Database
	// Checking Authentication Token
	if u.AuthToken == UKey {
		return true
	}
	return false
}
