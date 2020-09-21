package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerResponse(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open("static/index.html")
	p := make([]byte, 1024)
	if e == nil {
		f.Read(p)
		fmt.Fprint(w, string(p))
	} else {
		fmt.Fprint(w, "Hubo un error")
	}
}

func main() {
	http.HandleFunc("/", handlerResponse)
	fmt.Println(getPort())
	http.ListenAndServe(":"+getPort(), nil)
}

func getPort() string {
	fmt.Print(os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}
	return "80"
}
