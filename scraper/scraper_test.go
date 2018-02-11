package scraper

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

//TestInjectIdNormal make sure the id is being extracted correctly
func TestInjectIdNormal(t *testing.T) {
	offer := Offer{}
	mockDoc := goquery.Document{}
	err := offer.injectID(&mockDoc)
	if err != nil {
		t.Error(err)
	} else {

	}
}

//TestInjectIdMissing make sure it throws a specific error when id cannot be found
func TestInjectIdMissing(t *testing.T) {

}

func TestInjectTitleNormal(t *testing.T) {

}

func TestInjectPriceNormal(t *testing.T) {

}

func TestInjectPriceMissing(t *testing.T) {

}

func TestInjectAreaNormal(t *testing.T) {

}

func TestInjectAreaMissing(t *testing.T) {

}

func TestInjectAddressNormal(t *testing.T) {

}

func TestInjectAddressMissing(t *testing.T) {

}

func TestInjectDescriptionNormal(t *testing.T) {

}

func TestInjectDescriptionMissing(t *testing.T) {

}

func TestInjectLanguageNormal(t *testing.T) {

}

func TestInjectLanguageMissing(t *testing.T) {

}
