package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "failed to read body data error: %v\n", err)
		return
	}
	_, _ = fmt.Fprintf(w, "read body data: %s\n", string(bodyBytes))

	// 第二次再读取，不会报错，也不会读取到数据
	bodyBytes, err = io.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "failed to read body data error: %v\n", err)
		return
	}
	_, _ = fmt.Fprintf(w, "read body data: %s\n", string(bodyBytes))

}

func readQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for key, value := range query {
		_, _ = fmt.Fprintf(w, "key:%s, value%v", key, value)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/query", readQuery)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
