package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	);
	io.WriteString(res, `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				<h1>Hello World!</h1>
			</body>
		</html>
	`);
}

func main() {
	http.HandleFunc("/", index);
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	);
	http.ListenAndServe(":9000", nil);
}