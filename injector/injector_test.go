package injector

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

//TODO add another field to enter expected error(s)
type LiveInputAndAnswer struct {
	URL    string
	answer Offer
}

//TODO add another field to enter expected error(s)
type MockInputAndAnswer struct {
	FileName string
	answer   Offer
}

//test data
//This one is deactivated
var test1 = LiveInputAndAnswer{"https://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Prenzlauer-Berg.6584335.html", Offer{6584335, "3 Wochen Zwischenmiete. Ab Sofort", "", false, "Hufelandstr., 10407 Berlin Prenzlauer Berg", 0, 0, 0, 0, 0, 0, 0, 0, 0}}

//This one is deactivated
var test2 = LiveInputAndAnswer{"https://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Wedding.6566296.html", Offer{6566296, "schöne unmöbilierte WG-Zimmer(15 und 6m2) in 2er WG direkt am Leopoldplatz", "Simon Stracke", false, "Turiner Strasse, 13347 Berlin Wedding", 0, 0, 0, 0, 0, 0, 0, 0, 0}}

//This one is still active (last check: 2018.03.30)
var test3 = LiveInputAndAnswer{"https://www.wg-gesucht.de/en/wg-zimmer-in-Berlin-Tiergarten.4217243.html", Offer{4217243, "1 Zimmer, zentral, ruhig, möbl., 1. April 2018 ++ Anfragen nur telefonisch ++ requests only by telephone", "Vera Gardhoff", true, "Moabit, 10555 Berlin Tiergarten", 0, 0, 0, 0, 0, 0, 0, 0, 0}}

//TODO 1 for offer with no picture
//TODO 1 for offer with both from and to date
//TODO a whole bunch of offers with different sizes and prices
var liveTests = [...]LiveInputAndAnswer{test1, test2, test3}

//Part 1: Test each Inject function
//Goal: Make sure each function works, esp. when only 1 function get's updated

//Part 1.1: Test InjectActiveness
func TestInjectActiveness(t *testing.T) {
	for _, element := range liveTests {
		doc, err := goquery.NewDocument(element.URL)
		if err != nil {
			t.Fatalf("failed to get live data from '%s'", element.URL)
		}

		newOffer := Offer{}
		offer, err := InjectActiveness(&newOffer, doc)
		if err != nil {
			//TODO not necessarely: for invalid sites, there might be no id to be injected, and should throw an error
			t.Fatalf("inject activeness failed: '%s'", err)
		}
		if offer.IsActive() != element.answer.IsActive() {
			t.Errorf("activeness does not match: expecting '%t', got '%t' instead for '%s'", element.answer.IsActive(), offer.IsActive(), element.URL)
		}
	}
}

func TestInjectAdID(t *testing.T) {
	for _, element := range liveTests {
		doc, err := goquery.NewDocument(element.URL)
		if err != nil {
			t.Fatalf("failed to get live data from '%s'", element.URL)
		}

		newOffer := Offer{}
		offer, err := InjectAdID(&newOffer, doc)
		if err != nil {
			//TODO not necessarely: for invalid sites, there might be no id to be injected, and should throw an error
			t.Fatalf("inject id failed: '%s'", err)
		}
		if offer.AdID() != element.answer.AdID() {
			t.Errorf("AdID does not match: expecting %d, got %d instead", element.answer.AdID(), offer.AdID())
		}
	}
}

func TestInjectAdTitle(t *testing.T) {
	for _, element := range liveTests {
		doc, err := goquery.NewDocument(element.URL)
		if err != nil {
			t.Fatalf("failed to get live data from '%s'", element.URL)
		}

		newOffer := Offer{}
		offer, err := InjectAdTitle(&newOffer, doc)
		if err != nil {
			//TODO not necessarely: for invalid sites, there might be no id to be injected, and should throw an error
			t.Fatalf("inject title failed for '%s': '%s'", element.URL, err)
		}
		if offer.Title() != element.answer.Title() {
			t.Errorf("Title does not match: expecting '%s', got '%s' instead\n", element.answer.Title(), offer.Title())
		}
	}
}

func TestInjectOfferAddress(t *testing.T) {
	for _, element := range liveTests {
		doc, err := goquery.NewDocument(element.URL)
		if err != nil {
			t.Fatalf("failed to get live data from '%s'", element.URL)
		}

		newOffer := Offer{}
		offer, err := InjectAddress(newOffer, doc)
		if err != nil {
			//TODO not necessarely: for invalid sites, there might be no address to be injected, and should throw an error
			t.Fatalf("inject address failed for '%s': '%s'", element.URL, err)
		}
		if offer.address != element.answer.address {
			t.Errorf("Address does not match: expecting '%s', got '%s' instead\n", element.answer.address, offer.address)
		}
	}
}

//Part 2: Test with dynamically generated URLs (e.g. latest 1000 Ads)
//Goal: Make sure it does not crash or something

//Part 3: Benchmark with local data
//Goal: See if there is any noticable issue with each injector, and if it is necessary to use goroutine
