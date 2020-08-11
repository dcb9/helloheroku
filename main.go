package main

import "fmt"
import "net/http"
import "flag"


func main() {
	var bind = flag.String("bind", ":8080", "")
	flag.Parse()
	fmt.Printf("bind to %s\r\n", *bind)

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello heroku")
	})

	http.ListenAndServe(*bind, nil)
}

