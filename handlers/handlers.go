package ourcode

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		RenderWithError(w, "Method not allowed", 405)
		return
	}

	if r.URL.Path != "/" {
		RenderWithError(w, "Page not found", 404)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderWithError(w, "Template not found", 404)

		return
	}

	data := PageData{}
	if err := tmpl.Execute(w, data); err != nil {

		RenderWithError(w, "Internal server error", 500)

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

		RenderWithError(w, "Bad request", 400)

		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	// Validate input
	if text == "" {

		RenderWithError(w, "Text not found", 400)

		return

	}

	if banner == "" {
		banner = "standard"
	}

	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {

		RenderWithError(w, "Template not found", 400)

		return
	}

	result, err := GenerateASCIIArt(text, banner)
	if err != nil {
		RenderWithError(w, "Error generating ASCII art:", 500)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderWithError(w, "Template not found", 404)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		RenderWithError(w, "Internal server error", 500)
		return
	}
}
