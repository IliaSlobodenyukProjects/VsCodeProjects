package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

func main() {
	tmpl, err := template.New("example.html").Funcs(
		template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("example.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	todos := []Todo{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			param := r.FormValue("id")
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true
		}

		err := tmpl.Execute(w, todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8081", nil)
}
