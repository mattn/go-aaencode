package main

import (
	. "github.com/mattn/go-aaencode"
	. "http"
)

func main() {
	ListenAndServe(":8888", AAFilter(FileServer(Dir("static"))))
}
