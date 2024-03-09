package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getFoo(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("header", k, "value", v)
	}
	w.Write([]byte("bar"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("header", k, "value", v)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body", string(data))
	w.Write([]byte("login successful"))
}

func main() {
	http.HandleFunc("/foo", getFoo)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":1729", nil)
}
