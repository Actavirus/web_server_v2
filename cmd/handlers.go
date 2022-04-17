package main
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)
type User struct {
	Name string
	Nationality string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("html/home.page.tmpl")
		check(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.Nationality = r.Form.Get("entered_nationality")
		t, err := template.ParseFiles("html/foo.page.html")
		check(err)
		t.Execute(w, myUser)
	}
}

func fooPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Foopage")
}

func barPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Barpage")
}

