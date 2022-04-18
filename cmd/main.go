package main
import (
	"net/http"
	"log"
)

func main() {
	// Создайте мультиплексор для маршрутизации входящих запросов
	m := http.NewServeMux()
	// Все URL-адреса будут обрабатываться этой функцией
	m.HandleFunc("/", homePage)
	m.HandleFunc("/get", getPage)
	m.HandleFunc("/getJSON", getJSONPage)
	// Создайте сервер, прослушивающий порт 8000.
	log.Println("Starting a web server at http://127.0.0.1:8000")
	s := &http.Server{
		Addr: ":8000",
		Handler: m,
	}
	// Продолжайте обрабатывать новые запросы, пока не возникнет ошибка
	log.Fatal(s.ListenAndServe())
}