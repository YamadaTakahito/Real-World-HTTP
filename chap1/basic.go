package chap1

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func basic() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listeing :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
