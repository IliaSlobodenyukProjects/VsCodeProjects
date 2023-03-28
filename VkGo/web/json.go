package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func main() {
	todos := []Todo{
		{"Learn GO", false},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileContents, err := ioutil.ReadFile("example.html")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(fileContents)
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request", r.URL.Path)
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet:
			productsJson, _ := json.Marshal(todos)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(productsJson)
		case http.MethodPost:
			decoder := json.NewDecoder(r.Body)
			todo := Todo{}
			err := decoder.Decode(&todo)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			todos = append(todos, todo)
		case http.MethodPut:
			id := r.URL.Path[len("/todos/"):]
			index, _ := strconv.ParseInt(id, 10, 0)
			todos[index].Done = true
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8081", nil)
}
