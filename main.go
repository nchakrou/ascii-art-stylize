package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	ourcode "main.go/handlers"
)

func main() {
	// http.HandleFunc("/css",http.StripPrefix("/css",http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("templates/err.html")
		if r.URL.Path == "/css/" {
			temp.Execute(w, nil)
			return
		}
		http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/", ourcode.HomeHandler)
	http.HandleFunc("/ascii-art", ourcode.AsciiArtHandler)

	fmt.Println("Server starting on :8080")
	fmt.Println("Visit: http://127.0.0.1:8080")

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
