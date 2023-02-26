package services

import (
	"fmt"
	"log"
	"net/http"
)

func Execute() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Response status:", res.Status)
	fmt.Println("Response headers:", res.Header)
}
