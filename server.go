package main

import (
	"io"
	"net/http"
	"log"
	"time"
	"reflect"
	"fmt"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
    tstr := time.Now().Format("2006-01-02 03:04:05 PM")
    str := "hello, world!" + tstr
    fmt.Println(reflect.TypeOf(time.Now()));
	io.WriteString(w, str)
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Printf("server listen on port 1234");
	log.Fatal(http.ListenAndServe(":12345", nil))
}
