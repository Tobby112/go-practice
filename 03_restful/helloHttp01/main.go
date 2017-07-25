package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
	route       string
	handleCount int
}

func (hello *helloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Method == "GET" || req.Method == "POST" {
		fmt.Println(req.ContentLength)
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		hello.handleCount++
		w.Write([]byte(fmt.Sprintf("[%s] Hello, %s %s! by helloHandler(%d)", req.Method, firstname, lastname, hello.handleCount)))
	} else {
		http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Printf("Start\n")
	h := helloHandler{"/Hello", 0}
	http.Handle(h.route, &h)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	}
}
