package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/kr/pretty"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8181", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	resp, err := httputil.DumpRequest(r, true)
	pretty.Println(string(resp), err)
	w.WriteHeader(http.StatusOK)
}
