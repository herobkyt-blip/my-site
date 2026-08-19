package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
	<head><title>my first web site</title></head>
	<body>
	<h1>beginner programmer</h1>
	<p>i dont know what to write here, but i will figure it out</p>
	</body>
	</html>
`)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
	<head><title>About</title></head>
	<body>
	<h1>About Me</h1>
	<p>This is a simple about page.</p>
	</body>
	</html>
`)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
