package ourcode

import (
	"net/http"
	"text/template"
)

func renderWithError(w http.ResponseWriter, input, banner, errorMsg string) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  input,
		Banner: banner,
		Error:  errorMsg,
	}

	w.WriteHeader(http.StatusBadRequest)
	tmpl.Execute(w, data)
}
