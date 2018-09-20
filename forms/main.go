package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get current directory fail", err)
	}

	tmplFile := filepath.Join(cwd, "forms", "template", "layout.html")

	tmpl := template.Must(template.ParseFiles(tmplFile))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// Do something with details
		_ = details

		tmpl.Execute(w, struct{ Success bool }{Success: true})
	})

	http.ListenAndServe(":8888", nil)
}
