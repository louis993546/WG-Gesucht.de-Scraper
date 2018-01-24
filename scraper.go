package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//Price is what a price object can contain
type Price struct {
	Rent      int
	Utility   int
	Deposit   int
	Equipment int
	Other     int
}

//Gender of a person, right now it should only have 2 values: 0 for female, and 1 for male (get it?)
type Gender = int

//Images is just a slice of URLs, which basically should be treated like strings (because that's what they are)
type Images = []url.URL

//Date is just a Time, so that it looks slightly less confusing (Move in/out date 99.9% aren't that exact)
type Date = time.Time

//FurnishStatus have 3 possiblity: 0 for not furnished, 1 for partially furnished, 2 for furnished
type FurnishStatus = int

//Offer houses all the useful data that an offer can provide
type Offer struct {
	ID                int           //Ad ID of the offer
	Title             string        //Title of the offer
	Description       string        //Description of the offer, HTML text
	TotalSize         int           //Total house/apartment area in m^2
	RoomSize          int           //Toom area in m^2
	Address           string        //Address of the offer, might be accurate or not. TODO detect that.
	TotalRent         Price         //How much you have to pay per some unit of time
	AvailableFrom     Date          //The first date that you can move in
	AvailableTo       Date          //The last day that you have to move out
	Capacity          int           //Maximum human capaciity of the apartment/house
	OccupantsMale     int           //How many occupants are male
	OccupantsFemale   int           //How many occupants are female
	OccupantsAgeFrom  int           //Age of the youngest current occupant
	OccupantsAgeTo    int           //Age of the oldest current occupant
	TargetAgeFrom     int           //Minimum age of their target future roommate
	TargetAgeTo       int           //Maximum age of their target future roommate
	TargetGender      Gender        //What gender of people are they looking for
	Languages         []string      //TODO list of Languages (but should they be int or strings?)
	ImageUrls         Images        //Urls of those images
	Floor             int           //which floor the room/apartment is in. TODO it might be not int?
	HasElevator       bool          //True if the building has elevator
	HasWashingMachine bool          //True if the apartment/house/building has washing machine
	HasBalcony        bool          //True if apartment/house has a balcony
	Furnishes         FurnishStatus //TODO enum(y/n/partly)
	PosterName        string        //Name of the person who post this ad
}

//Request houses all the useful data that a request can provide
type Request struct {
	ID          int
	name        string
	description string
}

func main() {
	// thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/6507144.html")
	thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Wedding.4497377.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(thissssOffer)
}

func (offer *Offer) injectID(doc *goquery.Document) error {

}

func (offer *Offer) injectTitle(doc *goquery.Document) error {
	// doc.Find("#main_content").Find("#main_column").Find(".panel-body").Find(".headline.headline-detailed-view-title").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("title %d = %s\n", index, strings.TrimSpace(item.Text()))
	// })
}

//injectPrices extract the prices from doc and put it into offer. It returns an error if:
//1) Multiple conflicting prices found
//2) no prices found
//It will try to fetch multiple prices if possible
func (offer *Offer) injectPrices(doc *goquery.Document) error {
	// doc.Find("#graph_wrapper").Find(".basic_facts_bottom_part").Find(".amount").Each(func(index int, item *goquery.Selection) {
	// 	// fmt.Printf("rent = %s\n", strings.TrimSpace(item.Text()))
	// 	// *offer.price = int(item.Text())	//TODO somthing like this probably
	// })
}

//injectArea extract the the areas from doc and put it into offer. It returns an error if:
//1) Multiple conflicting area found
//2) Room area is larger than house area
//3) Room area not found
func (offer *Offer) injectArea(doc *goquery.Document) error {
	// doc.Find("#rent_wrapper").Find("label.amount").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("area[%d] = %s\n", index, strings.TrimSpace(item.Text()))
	// })
}

func (offer *Offer) injectAddress(doc *goquery.Document) error {
	// doc.Find("#main_content").Find("#main_column").Find(".panel-body").Find(".col-sm-4").Find("[onclick]").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("address = %s\n", strings.TrimSpace(item.Text()))
	// })
}

func (offer *Offer) injectDescription(doc *goquery.Document) error {
	// doc.Find("div.freitext").Each(func(index int, item *goquery.Selection) {
	// 	//fmt.Printf("Description:\n%s\n", strings.TrimSpace(item.Text()))
	// 	thisOffer.description = strings.TrimSpace(item.Text())
	// })
}

func (offer *Offer) injectLanguages(doc *goquery.Document) error {
	// doc.Find("img.flgS").Each(func(index int, item *goquery.Selection) {
	// 	title, _ := item.Attr("title")
	// 	fmt.Printf("flag %d = %s\n", index, title)
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

	err = thisOffer.injectPrices(doc)
	if err != nil {
		return thisOffer, err
	}
	err = thisOffer.injectArea(doc)
	if err != nil {
		return thisOffer, err
	}
	err = thisOffer.injectAddress(doc)
	if err != nil {
		return thisOffer, err
	}
	err = thisOffer.injectDescription(doc)
	if err != nil {
		return thisOffer, err
	}
	err = thisOffer.injectTitle(doc)
	if err != nil {
		return thisOffer, err
	}

	return thisOffer, nil
}
