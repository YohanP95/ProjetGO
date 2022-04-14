package main

import "net/http"

type Task struct {
    Description string
    Done        bool
}

var task []Task

func main () {
    http.ListenAndServe(":8080", nil)
}
