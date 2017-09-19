package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func mainPageView(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "This is main page of site.") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./src/site/login.gptl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		getUserLogin := r.Form["username"]
		getUserPassword := r.Form["password"]
		if (getUserLogin[0] == "admin") && (getUserPassword[0] == "admin") {
			fmt.Fprintln(w, "Fucking right.")
		} else{
			fmt.Fprintln(w, "Wrong password.")
		}
	}
}

func main() {
	http.HandleFunc("/", mainPageView) // setting router rule
	http.HandleFunc("/admin", login)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}