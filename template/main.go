package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get current directory fail", err)
	}

	// specific file path
	tmplFile := filepath.Join(cwd, "template", "tmplfile", "layout.html")

	tmpl := template.Must(template.ParseFiles(tmplFile))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
				{Title: "Task 4", Done: false},
			},
		}

		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8888", nil)
}
