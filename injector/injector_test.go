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

	test1 := InputAndAnswer{"https://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Prenzlauer-Berg.6584335.html", Offer{6584335, "", "", "", 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	doc, err := goquery.NewDocument(test1.URL)
	if err != nil {
		t.Fatalf("failed to get live data from '%s'", test1.URL)
	}

	newOffer := Offer{}
	offer, err := InjectAdID(&newOffer, doc)
	if err != nil {
		t.Fatalf("inject id failed: '%s'", err)
	}
	if offer.AdID() != test1.answer.AdID() {
		t.Errorf("AdID does not match: expecting %d, got %d instead", test1.answer.AdID(), offer.AdID())
	}

	//If i am expecting error, check if offer.ID is empty
	//If i am expecting id, check offer.ID
}
