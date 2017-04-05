package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/codeNut247/golang_website/Models"
)

type pageData struct {
	Title     string
	FirstName string
	User      Models.User
}

var pd pageData

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/signin", signin)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8080", nil)
}

func sendView(w http.ResponseWriter, tempName string, pd pageData) {
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, tempName, pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func signin(w http.ResponseWriter, req *http.Request) {
	//pd = pageData{Title: "Sign Up"}
	pd.Title = "Sign In"
	if req.Method == http.MethodPost {
		pd.User.UserName = req.FormValue("username")
		pd.User.Email = req.FormValue("email")
		pd.User.Password = req.FormValue("password")
		if req.FormValue("remember") == "on" {
			pd.User.Remember = true
		}
		pd.User.IsLogedIn = true
		log.Println(pd.User)
		log.Println(req.FormValue("remember"))
	}
	/*
		pd.Title = "Sign Up"
		if req.Method == http.MethodPost {
			userForm := new(Models.User)
			errs := binding.Bind(req, userForm)
			if errs.Handle(w) {
				return
			}
			pd.User.UserName = userForm.UserName
			log.Println(userForm.Message)
			log.Println(userForm.UserName)
		}
	*/
	sendView(w, "signinCreate.gohtml", pd)
}

func index(w http.ResponseWriter, req *http.Request) {
	//pd = pageData{Title: "Index"}
	pd.Title = "Index"
	sendView(w, "templist.gohtml", pd)
}

func about(w http.ResponseWriter, req *http.Request) {
	//pd = pageData{Title: "About"}
	pd.Title = "About"
	sendView(w, "templist.gohtml", pd)
}

func contact(w http.ResponseWriter, req *http.Request) {
	//pd = pageData{Title: "Contact"}
	pd.Title = "Contact"
	sendView(w, "templist.gohtml", pd)
}

func apply(w http.ResponseWriter, req *http.Request) {
	//pd = pageData{Title: "Apply"}
	pd.Title = "Apply"
	if req.Method == http.MethodPost {
		pd.FirstName = req.FormValue("fname")
	}
	sendView(w, "createview.gohtml", pd)
}
