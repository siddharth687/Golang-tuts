package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to URLS in Golang")
	fmt.Println(myurl)
	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["v"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=hitesh",
	}
	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
