package main
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"io/ioutil"
	"net/url"
	"encoding/json"
)

const apiURL = "https://api.stackexchange.com/2.2/posts?"

// Structs for JSON decoding
type postItem struct {
	Score int `json:"score"`
	Link string `json:"link"`
}

type postsType struct {
	Items []postItem `json: "items"`
}

type User struct {
	Name string
	Nationality string
}



func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// https://learntutorials.net/ru/go/topic/756/http-%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80
func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("ui/html/home.page.html")
		check(err)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		myUser := User{}
		myUser.Name = r.Form.Get("entered_name")
		myUser.Nationality = r.Form.Get("entered_nationality")
		t, err := template.ParseFiles("ui/html/greetings.page.html")
		check(err)
		t.Execute(w, myUser)
	}
}


// https://learntutorials.net/ru/go/topic/1422/http-%d0%ba%d0%bb%d0%b8%d0%b5%d0%bd%d1%82
func getPage(w http.ResponseWriter, r *http.Request) {
	// Основной GET
	// Выполните основной запрос GET и распечатайте содержимое сайта (HTML).
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	// It is important to defer resp.Body.Close(), else resource leaks will occur.
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Will print site contents (HTML) to output.
	fmt.Println(string(data))
	w.Write([]byte(string(data)))
}

// https://learntutorials.net/ru/go/topic/1422/http-%d0%ba%d0%bb%d0%b8%d0%b5%d0%bd%d1%82
// GET с параметрами URL и ответом JSON
// Запрос на 10 самых последних активных сообщений StackOverflow, использующих API-интерфейс Stack Exchange.
func getJSONPage(w http.ResponseWriter, r *http.Request) {
	// Set URL parameters on declaration
	values := url.Values{
		"order": []string{"desc"},
		"sort": []string{"activity"},
		"site": []string{"stackoverflow"},
	}

	// URL parameters can also be programmatically set
	values.Set("page", "1")
	values.Set("pagesize", "10")

	resp, err := http.Get(apiURL + values.Encode())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// To compare status codes, you should always use the status constants
	// provided by the http package.
	if resp.StatusCode != http.StatusOK {
		panic("Request was not OK: " + resp.Status)
	}

	// Example of JSON decoding on a reader.
	dec := json.NewDecoder(resp.Body)
	var p postsType
	err = dec.Decode(&p)
	if err != nil {
		panic(err)
	}

	fmt.Println("Top 10 most recently active StackOver posts: ")
	fmt.Println("Score", "Link")
	for _, post := range p.Items {
		fmt.Println(post.Score, post.Link)
		w.Write([]byte(post.Link))
	}
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Test of protocol HTTP", nil)
	t, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}