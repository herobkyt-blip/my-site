package main

import (
	"fmt"
	"net/http"
	"os"
)

const pageStyle = `
<style>
	body {
		font-family: sans-serif;
		background-color: #1a1a1a;
		color: #f0f0f0;
		max-width: 700px;
		margin: 40px auto;
		padding: 0 20px;
	}
	a {
		color: #4ea1f3;
	}
</style>
`

const navBar = `
<nav>
	<a href="/">Home</a> | <a href="/projects">Projects</a>
</nav>
`

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
	<head>
	<title>my first web site</title>
	%s
	</head>
	<body>
	%s
	<h1>beginner programmer</h1>
	<p>i dont know what to write here, but i will figure it out</p>
	<p><a href="https://t.me/netlydeidlyakadra" target="_blank">my telegramm channel</a></p>
	</body>
	</html>
`, pageStyle, navBar)
}

func projectPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<html>
	<head>
	<title>Projects</title>
	%s
	</head>
	<body>
	%s
	<h1>it's my projects</h1>
	</style>
	</head>
	<body>
	<h1>it's my projects</h1>
	<p><a href="https://t.me/dynastynalbot" target="_blank">Dynasty Bot</a> — Telegram bot on the Go. For Telegram shop.</p>
	<a href="/">Home</a>
	<p><a href="https://t.me/netlydeidlyakadra" target="_blank">my telegramm channel</a></p>
	</body>
	</html>
`, pageStyle, navBar)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/projects", projectPage)
	fmt.Println("Сервер запущен на http://localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}
