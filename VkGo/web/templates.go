package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	tmpl := template.New("main")
	tmpl, _ := tmpl.Parse("htmlstring")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		c := http.Client{}
		resp, err := c.Get("/something" + path)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		tmpl.Execute(w, string(body))
	})

	http.ListenAndServe(":8081", nil)
}
