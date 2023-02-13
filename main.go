package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var c http.Client
	os.Setenv("HTTP_PROXY", "notexist.fakedomain123fooabc.com")
	res, err := c.Get("http://example.com")
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)
}
