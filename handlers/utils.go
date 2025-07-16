package ourcode

import (
	"log"
	"net/http"
	"text/template"
)

func RenderWithError(w http.ResponseWriter, errorMsg string, errNumb int) {
	errstr := Initialiseerr(errNumb, errorMsg)
	w.WriteHeader(errNumb)
	errtempl, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}
	if err := errtempl.Execute(w, errstr); err != nil {
		// If template execution fails, fallback to plain error
		http.Error(w, "Request could not be processed", http.StatusInternalServerError)
	}
}
