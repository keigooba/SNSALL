package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello world")
}
