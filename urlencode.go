// urlencode converts a single argument (a raw url string without a query) to an escaped version.
// usage: urlencode rawurl
package main

import (
	"fmt"
	"log"
	"os"
	"net/url"
)

// urlencode implements the underlying string to string conversion
func urlencode(rawurl string) string {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}
	return u.String()
}

// main expects a single argument (rawurl) and prints the escaped version
func main() {
	fmt.Println(urlencode(os.Args[1]))
}
