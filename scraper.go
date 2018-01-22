package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

//Offer houses all the useful data that an offer can provide
type Offer struct {
	id          int
	title       string
	description string
	houseSize   int
	roomSize    int
	address     string
	price       int
}

//Request houses all the useful data that a request can provide
type Request struct {
	id          int
	name        string
	description string
}

func main() {
	thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/6507144.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(thissssOffer)

	// doc, err := goquery.NewDocument("http://www.wg-gesucht.de/6507144.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// thisOffer := Offer{}

	// //title
	// doc.Find("#main_content").Find("#main_column").Find(".panel-body").Find(".headline.headline-detailed-view-title").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("title %d = %s\n", index, strings.TrimSpace(item.Text()))
	// })

	// //languages
	// doc.Find("img.flgS").Each(func(index int, item *goquery.Selection) {
	// 	title, _ := item.Attr("title")
	// 	fmt.Printf("flag %d = %s\n", index, title)
	// })

	// //description
	// doc.Find("div.freitext").Each(func(index int, item *goquery.Selection) {
	// 	//fmt.Printf("Description:\n%s\n", strings.TrimSpace(item.Text()))
	// 	thisOffer.description = strings.TrimSpace(item.Text())
	// })

	// //areas
	// doc.Find("#rent_wrapper").Find("label.amount").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("area[%d] = %s\n", index, strings.TrimSpace(item.Text()))
	// })

	// //price
	// doc.Find("#graph_wrapper").Find(".basic_facts_bottom_part").Find(".amount").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("rent = %s\n", strings.TrimSpace(item.Text()))
	// })

	// //address
	// doc.Find("#main_content").Find("#main_column").Find(".panel-body").Find(".col-sm-4").Find("[onclick]").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("address = %s\n", strings.TrimSpace(item.Text()))
	// })
}

//ScrapRequest turns an offer url to an Offer struct
//TODO url formatter: check if the url is from wg first
func ScrapRequest(url string) (offer Offer, err error) {
	var thisOffer Offer
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return thisOffer, err
	}
	fmt.Println(doc)
	//TODO in progress

	return thisOffer, nil
}
