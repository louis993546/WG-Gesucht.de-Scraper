package validator

import (
	"regexp"
)

//LooksWgGesucht checks if at the very least the url looks like it is from wg-gesucht.de
//i.e. does the url starts with "www.wg-gesucht.de" (+- http(s)://)
func LooksWgGesucht(url string) bool {
	regex := regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/`)
	return regex.MatchString(url)
}

//IsOfferList checks if the url is a list of offers on wg-gesucht.de
func IsOfferList(url string) bool {
	return false
}

//IsRequestList checks if the url is a list of requests on wg-gesucht.de
func IsRequestList(url string) bool {
	regex := regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/(.*-gesucht)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)
	return regex.MatchString(url)
}

//IsEnglish checks if the url is an english version of some page on wg-gesucht.de
func IsEnglish(url string) bool {
	regex := regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/en/`)
	return regex.MatchString(url)
}

//IsSpanish checks if the url is a spanish version of some page on wg-gesucht.de
func IsSpanish(url string) bool {
	regex := regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/es/`)
	return regex.MatchString(url)
}

//IsGerman checks if the url is a german version of some page on wg-gesucht.de
func IsGerman(url string) bool {
	return LooksWgGesucht(url) //because the german one does not have any identifier
}
