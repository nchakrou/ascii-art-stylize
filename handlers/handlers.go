package ourcode

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//  fmt.Fprintln(w, "<h1>Your welcome brother</h1>")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data := PageData{}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate input
	if text == "" {
		renderWithError(w, text, banner, "/ Text is required")
		return
	}

	if banner == "" {
		banner = "standard"
	}

	// Validate banner type
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		renderWithError(w, text, banner, " / Invalid banner type")
		return
	}

	// Check for invalid characters
	for _, char := range text {
		if char < 32 || char > 126 {
			if char != '\n' && char != '\r' {
				renderWithError(w, text, banner, " / Text contains unsupported characters")
				return
			}
		}
	}

	// Generate ASCII art
	result, err := GenerateASCIIArt(text, banner)
	if err != nil {
		renderWithError(w, text, banner, "     / Error generating ASCII art: "+err.Error())
		return
	}

	// Render result
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
