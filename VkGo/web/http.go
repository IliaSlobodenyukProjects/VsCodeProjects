package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		w.WriteHeader(200)
		w.Write(body)
	})

	http.ListenAndServe(":8081", nil)
}
