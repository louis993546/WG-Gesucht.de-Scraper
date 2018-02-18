package scraper

import (
	"encoding/json"
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
	output, _ := json.Marshal(price)
	return string(output)
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

func (oa OfferAddress) String() string {
	output, _ := json.Marshal(oa)
	return string(output)
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
	output, _ := json.Marshal(ns)
	return string(output)
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
	output, _ := json.Marshal(offer)
	return string(output)
}

//Request houses all the useful data that a request can provide
type Request struct {
	ID          int
	Name        string
	Description string
}

func (request Request) String() string {
	output, _ := json.Marshal(request)
	return string(output)
}

func main() {
	// thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/6507144.html")
	thissssOffer, err := ScrapRequest("http://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Wedding.4497377.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(thissssOffer)
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

	//technically all of the following can be done with goroutine -> need to figure out how to
	//1. kill all other routines when fatal in one of them
	//2. terminate once all the goroutines finished

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
		//TODO this should not return (because images should be nullable)
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
