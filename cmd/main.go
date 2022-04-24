package main

import (
	"log"
	"net/http"
)

func main() {
	// Создайте мультиплексор для маршрутизации входящих запросов
	m := http.NewServeMux()
	// Все URL-адреса будут обрабатываться этой функцией
	m.HandleFunc("/", homePage)
	m.HandleFunc("/vvod", vvodPage)
	m.HandleFunc("/get", getPage)
	m.HandleFunc("/getJSON", getJSONPage)
	m.HandleFunc("/servehttp", serveHTTP)
	m.HandleFunc("/zadacha1", zadacha1)

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// Используем функцию mux.Handle() для регистрации обработчика для
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	m.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Создайте сервер, прослушивающий порт 8000.
	log.Println("Starting a web server at http://127.0.0.1:8000")
	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}
	// Продолжайте обрабатывать новые запросы, пока не возникнет ошибка
	log.Fatal(s.ListenAndServe())
}
