//app1

package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/", test)

	fmt.Println("Service Started ...")
	http.ListenAndServe(":15001", nil)

}


func test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello !! test is ok"))

}
