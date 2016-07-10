package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

// Comment has a comment and a creation timestamp
type Comment struct {
	comment   string
	timestamp int64
}

// Annotation has an annotation and a creation timestamp
type Annotation struct {
	annotation string
	timestamp  int64
}

func comment(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%q\n", dump)

	decoder := json.NewDecoder(r.Body)
	var comment Comment
	err = decoder.Decode(&comment)
	if err != nil {
		http.Error(w, "{\"response\":\"error\"}", http.StatusBadRequest)
		return
	}
	io.WriteString(w, "{\"response\":\"success\"}")
}

func annotation(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%q\n", dump)

	decoder := json.NewDecoder(r.Body)
	var annotation Annotation
	err = decoder.Decode(&annotation)
	if err != nil {
		http.Error(w, "{\"response\":\"error\"}", http.StatusBadRequest)
		return
	}
	io.WriteString(w, "{\"response\":\"success\"}")
}

func main() {
	http.HandleFunc("/comment/", comment)
	http.HandleFunc("/annotation/", annotation)
	http.ListenAndServe(":9090", nil)
}
