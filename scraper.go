package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//Price is what a price object can contain. Note that the total is not always the sum of all of the values, and I am still trying to figure out how that should work
type Price struct {
	Rent      int
	Utility   int
	Deposit   int
	Equipment int
	Other     int
}

//turn struct into json for log purposes. DO NOT use it for the actual output!
func (price Price) String() string {
	return "{'rent': " + string(price.Rent) + ", 'utility;: " + string(price.Utility) + ", 'deposit;: " + string(price.Deposit) + ", 'equipment;: " + string(price.Equipment) + ", 'other;: " + string(price.Other) + "}"
}

const (
	//GenderFemale means gender = female, pretty self explanatory
	GenderFemale = 0

	//GenderMale means gender = male, pretty self explanatory
	GenderMale = 1

	//GenderAny means it does not matter, mostly for filtering
	GenderAny = 100
)

//Gender of a person, can be Male, Female, or Any (i.e. I don't care)
type Gender = int

//Images is just a slice of URLs, which basically should be treated like strings (because that's what they are)
type Images = []url.URL

//Date is just a Time, so that it looks slightly less confusing (Move in/out date 99.9% aren't that exact)
type Date = time.Time

const (
	//FurnishedNo means the room/apartment is not furnished
	FurnishedNo = 0

	//FurnishedPartially means the room/apartment is only partically furnished
	FurnishedPartially = 1

	//FurnishedYes means the room/apartment is fully furnished.
	FurnishedYes = 2
)

//FurnishStatus have 3 possiblity: 0 for not furnished, 1 for partially furnished, 2 for furnished
type FurnishStatus = int

//OfferAddress house 2 values, so that in the future filtering is a bit easier
//future TODO: lat & long (i.e. ask Google where that address is actually)
type OfferAddress struct {
	AddressString string
	IsApproximate bool
	Latitude      float32
	Longitude     float32
}

//TODO figure out how to turn boolean to string
func (oa OfferAddress) String() string {
	return "{'address_string': " + oa.AddressString + ", 'is_approximate': " + string(oa.IsApproximate) + "}"
}

//NetworkSpeed has min and max. This exist because how WG-Gesucht works sucks
//"Slower than 1Mbit/s"		-> 0 <-> 1
//"Up to 3Mbit/s" 			-> 0 <-> 3
//"Faster than 50Mbit/s"	-> 50 <-> NetworkSpeedInfinite
//You get the idea.
type NetworkSpeed struct {
	Min int
	Max int
}

const (
	//NetworkSpeedInfinite means the network speed is unlimitedly or unspecifiecly high, e.g. when network speed is "Faster than 50Mbit/s", the max should be -1
	NetworkSpeedInfinite = -1
)

func (ns NetworkSpeed) String() string {
	return "{'min': " + string(ns.Min) + ", 'max': " + string(ns.Max) + "}"
}

//YearMonth is another wrapper around time.Time: it's for thing that require even less precision than Date, i.e. "Member since"
type YearMonth = time.Time

//Offer houses all the useful data that an offer can provide
//TODO there are multiple types of offer: WG, whole apartment, etc, and they will have very different fields, but they are all "Offer"s. Need to re-write this struct (kinda).
type Offer struct {
	ID                 int           //Ad ID of the offer
	Title              string        //Title of the offer
	Description        string        //Description of the offer, HTML text. TODO see if it is possible to split it into it's 4 category
	TotalSize          int           //Total house/apartment area in m^2
	RoomSize           int           //Toom area in m^2
	Address            OfferAddress  //Address of the offer, might be accurate or not. TODO detect that.
	Rent               Price         //How much you have to pay per some unit of time
	AvailableFrom      Date          //The first date that you can move in
	AvailableTo        Date          //The last day that you have to move out
	Capacity           int           //Maximum human capaciity of the apartment/house
	OccupantsMale      int           //How many occupants are male
	OccupantsFemale    int           //How many occupants are female
	OccupantsAgeMin    int           //Age of the youngest current occupant
	OccupantsAgeMax    int           //Age of the oldest current occupant
	Languages          []string      //TODO list of Languages (but should they be int or strings?)
	TargetAgeMin       int           //Minimum age of their target future roommate
	TargetAgeMax       int           //Maximum age of their target future roommate
	TargetGender       Gender        //What gender of people are they looking for
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
	memberSince        YearMonth
}

func (offer Offer) String() string {
	return "{" + "}"
}

//Request houses all the useful data that a request can provide
type Request struct {
	ID          int
	Name        string
	Description string
}

func (request Request) String() string {
	return "{'id': " + string(request.ID) + ", 'name': '" + request.Name + "', 'description': '" + request.Description + "'}"
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

//Extract unique identifier from the page, in this case (WG-gesucht), it's the Ad Id
func (offer *Offer) injectID(doc *goquery.Document) error {
	//several more layers
	//	class="col-md-4"
	//		class="row"
	//			class="col-md-12"
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

//injectAddress extracts all the address stuff from the page:
//1. Address in text
//2. Is the address just an approximation
//3. Latitude (from onClick loagGMap)
//4. Longitude (from onClick loagGMap)
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

//injectAvailability extract the start date and end date of the offer into AvailableFrom and AvailableTo
func (offer *Offer) injectAvailability(doc *goquery.Document) error {

}

//injectCurrentOccupantSize extracts the info of the current occupant(s), i.e. OccupantsMale, and OccupantsFemale
func (offer *Offer) injectCurrentOccupantSize(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

//injectCurrentOccupantSize extracts the info of the current occupant(s), i.e. OccupantsAgeMin, and OccupantsAgeMax
func (offer *Offer) injectCurrentOccupantAge(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

//injectTargetLimition extracts
func (offer *Offer) injectTargetLimition(doc *goquery.Document) error {
	//ul-detailed-view-datasheet print_text_left
}

//injectLanguages extracts the list of languages that the offer has listed as the languages that he/she/they speak(s)
//TODO make a map that turns flag title back to standard language code. Also, check if golang has standard language code
func (offer *Offer) injectLanguages(doc *goquery.Document) error {
	// doc.Find("img.flgS").Each(func(index int, item *goquery.Selection) {
	// 	title, _ := item.Attr("title")
	// 	fmt.Printf("flag %d = %s\n", index, title)
	// })
}

//injectImages extracts the list of images urls (if available)
func (offer *Offer) injectImages(doc *goquery.Document) error {

}

//injectMinorDetails basically runs through the 3 column grid and put them into
//their corresponding field in the offer struct
func (offer *Offer) injectMinorDetails(doc *goquery.Document) error {
	//pretty sure the .Find() does not that multiple class that way
	doc.Find(".col-xs-6 .col-sm-4 .col-md-4 .print_text_left").Each(func(index int, item *goquery.Selection) {
		classes, _ := item.Find("span").First().Attr("class") //this should give you "glyphicons glyphicons-?????? noprint"
		text = item.Text()
		switch classes {
		case "glyphicons glyphicons-folder-closed noprint": //e.g. "Washing machine, Balcony, Basement/Cellar, Elevator, Pets are welcome "
			//-> washing machine
			//-> elevator
			//-> balcony
			//-> pet
			//-> basement
		case "glyphicons glyphicons-car noprint": //e.g. "Many free parking lots"
		case "glyphicons glyphicons-fire noprint": //e.g. "Central heating"
			//-> heating
		case "glyphicons glyphicons-fabric noprint": //e.g. "Polished floorboards"
		case "glyphicons glyphicons-wifi-alt noprint": //e.g. "DSL, WLAN 26-50 Mbit/s"
			//-> internet
		case "glyphicons glyphicons-bath-bathtub noprint": //e.g. "Bathtub"
			//-> bathtub
		case "glyphicons glyphicons-bed noprint": //e.g. " 3rd floor, furnished " (Yes, extra spaces included -_-)
			//-> floor
			//-> furnished

			//1. string to []string by commar
			//2. trim
			//3. regex "??? floor" or other form of levels
			//4. furnished, partically furnished, or not furnished
		case "glyphicons glyphicons-building noprint": //e.g. "Industrial building"
			//might want to put default here to catch other stuff
		}

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

//injectPosterName extract the (user)name of the person who post this Offer/Request, i.e. the "Name: " field
func (offer *Offer) injectPosterName(doc *goquery.Document) error {

}

//ScrapRequest turns an offer url to an Offer struct
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

	//technically all of the following can be done with goroutine

	err = thisOffer.injectID(doc)
	if err != nil {
		return thisOffer, err
	}

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
