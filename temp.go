package main

import (
	"fmt"
	"log"
)

func main() {
	// thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/6507144.html")
	thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Wedding.4497377.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(thissssOffer)
}
