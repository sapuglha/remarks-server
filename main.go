package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

// Comment has a message and a creation timestamp
type Comment struct {
	comment   string
	timestamp int64
}

func message(w http.ResponseWriter, r *http.Request) {
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
		io.WriteString(w, "{\"response\":\"invalid\"}")
		return
	}
	io.WriteString(w, "{\"response\":\"success\"}")
}

func main() {
	http.HandleFunc("/comment/", message)
	http.ListenAndServe(":9090", nil)
}
