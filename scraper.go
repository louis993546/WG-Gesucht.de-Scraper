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

//OfferAddress house 2 values, so that in the future filtering is a bit easier
//future TODO: lat & long (i.e. ask Google where that address is actually)
type OfferAddress struct {
	AddressString string
	IsApproximate bool
}

//NetworkSpeed has min and max. This exist because how WG-Gesucht works sucks
//"Slower than 1Mbit/s"		-> 0 - 1
//"Up to 3Mbit/s" 			-> 0 - 3
//"Faster than 50Mbit/s"	-> 50 - (some theoretical max value)
type NetworkSpeed struct {
	Min int
	Max int
}

//Offer houses all the useful data that an offer can provide
type Offer struct {
	ID                 int           //Ad ID of the offer
	Title              string        //Title of the offer
	Description        string        //Description of the offer, HTML text. TODO see if it is possible to split it into it's 4 category
	TotalSize          int           //Total house/apartment area in m^2
	RoomSize           int           //Toom area in m^2
	Address            OfferAddress  //Address of the offer, might be accurate or not. TODO detect that.
	TotalRent          Price         //How much you have to pay per some unit of time
	AvailableFrom      Date          //The first date that you can move in
	AvailableTo        Date          //The last day that you have to move out
	Capacity           int           //Maximum human capaciity of the apartment/house
	OccupantsMale      int           //How many occupants are male
	OccupantsFemale    int           //How many occupants are female
	OccupantsAgeFrom   int           //Age of the youngest current occupant
	OccupantsAgeTo     int           //Age of the oldest current occupant
	TargetAgeFrom      int           //Minimum age of their target future roommate
	TargetAgeTo        int           //Maximum age of their target future roommate
	TargetGender       Gender        //What gender of people are they looking for
	Languages          []string      //TODO list of Languages (but should they be int or strings?)
	ImageUrls          Images        //Urls of those images
	Floor              int           //which floor the room/apartment is in. TODO it might be not int?
	HasElevator        bool          //True if the building has elevator
	HasWashingMachine  bool          //True if the apartment/house/building has washing machine
	HasBalcony         bool          //True if apartment/house has a balcony
	Furnishes          FurnishStatus //TODO enum(y/n/partly)
	PosterName         string        //Name of the person who post this ad
	BuildingType       int           //There are 9 building types
	WheelchairFriendly bool          //Call "Handicapped Accessible" on the site
	Smokeable          int           //There are 4 possible values
	FlatshareType      []string      //Honestly I am still not sure what is is about
	Flooring           []string      //What material(s) is/are being used in the apartment's floor
	HeatingType        int           //There are 6 types of heating (excluding the one in flooring -_-)
	InternetSpeed      NetworkSpeed  //in Mb/s
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

//URLCheck at least make sure url is from WG-Gesucht.de. Also it make sure the site is in English mode
func URLCheck(url string) (cleanURL string, err error) {
	//offer & request format: http://www.wg-gesucht.de/{string}{7 digit id}.html
}

func (offer *Offer) injectID(doc *goquery.Document) error {

}

func (offer *Offer) injectTitle(doc *goquery.Document) error {
	// doc.Find("#main_content").Find("#main_column").Find(".panel-body").Find(".headline.headline-detailed-view-title").Each(func(index int, item *goquery.Selection) {
	// 	fmt.Printf("title %d = %s\n", index, strings.TrimSpace(item.Text()))
	// })
}

func (offer *Offer) injectDescription(doc *goquery.Document) error {
	// doc.Find("div.freitext").Each(func(index int, item *goquery.Selection) {
	// 	//fmt.Printf("Description:\n%s\n", strings.TrimSpace(item.Text()))
	// 	thisOffer.description = strings.TrimSpace(item.Text())
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

//injectPrices extract the prices from doc and put it into offer. It returns an error if:
//1) Multiple conflicting prices found
//2) no prices found
//It will try to fetch all the sub-prices if possible
func (offer *Offer) injectPrices(doc *goquery.Document) error {
	// doc.Find("#graph_wrapper").Find(".basic_facts_bottom_part").Find(".amount").Each(func(index int, item *goquery.Selection) {
	// 	// fmt.Printf("rent = %s\n", strings.TrimSpace(item.Text()))
	// 	// *offer.price = int(item.Text())	//TODO somthing like this probably
	// })
}

func (offer *Offer) injectAvailability(doc *goquery.Document) error {

}

func (offer *Offer) injectCurrentOccupantSize(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

func (offer *Offer) injectCurrentOccupantAge(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

func (offer *Offer) injectTargetLimition(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

func (offer *Offer) injectLanguages(doc *goquery.Document) error {
	// doc.Find("img.flgS").Each(func(index int, item *goquery.Selection) {
	// 	title, _ := item.Attr("title")
	// 	fmt.Printf("flag %d = %s\n", index, title)
	// })
}

func (offer *Offer) injectImages(doc *goquery.Document) error {

}

//injectMinorDetails basically runs through the 3 column grid and put them into
//their corresponding field in the offer struct
func (offer *Offer) injectMinorDetails(doc *goquery.Document) error {
	//pretty sure the .Find() does not that multiple class that way
	doc.Find(".col-xs-6 .col-sm-4 .col-md-4 .print_text_left").Each(func(index int, item *goquery.Selection) {
		detailType := item.Find("span").Attr("class") //this should give you "glyphicons glyphicons-?????? noprint"
		//floor
		//elevator
		//washing machine
		//balcony
		//furnished
		//heating
		//internet
		//bathtub
		//basement
		//pet
		//heating
		//telephone
		//flooring
		//tv
		//dishwasher
		//terrace
		//garden
		//shared garden
		//bicycle storage
		//green power
	})
	return nil
}

func (offer *Offer) injectPosterName(doc *goquery.Document) error {

}

//ScrapRequest turns an offer url to an Offer struct
//TODO url formatter: check if the url is from wg first
func ScrapRequest(url string) (offer Offer, err error) {
	var thisOffer Offer

	cleanURL, err := UrlCheck(url)
	if err != nil {
		return thisOffer, err
	}

	doc, err := goquery.NewDocument(cleanUrl)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectID(doc)

	err = thisOffer.injectTitle(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectDescription(doc)
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

	err = thisOffer.injectPrices(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectAvailability(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectCurrentOccupantSize(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectCurrentOccupantAge(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectTargetLimition(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectLanguages(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectImages(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectMinorDetails(doc)
	if err != nil {
		return thisOffer, err
	}

	err = thisOffer.injectPosterName(doc)
	if err != nil {
		return thisOffer, err
	}

	return thisOffer, nil
}
