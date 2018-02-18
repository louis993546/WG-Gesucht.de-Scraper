package injector

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/louistsaitszho/wggesuchtscraper/scraper"
)

//Extract unique identifier from the page, in this case (WG-gesucht), it's the Ad Id
func (offer *scraper.Offer) injectID(doc *goquery.Document) error {
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
