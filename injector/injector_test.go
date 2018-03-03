package injector

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type InputAndAnswer struct {
	URL    string
	answer Offer
}

//Part 1: Test with mock data
func TestWithMockData(t *testing.T) {
	//TODO loop through a list of mock document and their model answers
}

//Part 2: Test with live data
func TestWithLiveData(t *testing.T) {
	//TODO use a list of live websites and their expected results
	url := "https://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Kreuzberg.6475694.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		t.Fatalf("failed to get live data from '%s'", url)
	}

	newOffer := Offer{}
	offer, err := InjectAdID(&newOffer, doc)
	if err != nil {
		t.Fatalf("inject id failed: '%s'", err)
	}
	println(offer)

	//If i am expecting error, check if offer.ID is empty
	//If i am expecting id, check offer.ID
}
