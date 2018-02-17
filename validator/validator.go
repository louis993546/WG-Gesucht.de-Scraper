package validator

import (
	"regexp"
)

//RegexpWgGesuchtBase is the compiled regex for wg gesucht base
var RegexpWgGesuchtBase = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/`)

//RegexpWgGesuchtList is the compiled regex for wg gesucht list
var RegexpWgGesuchtList = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/(.*)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtRequestList is the compiled regex for wg gesucht request list
var RegexpWgGesuchtRequestList = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/(.*-gesucht)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtEnglish is the compiled regex for wg gesucht english site
var RegexpWgGesuchtEnglish = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/en/`)

//RegexpWgGesuchtSpanish is the compiled regex for wg gesucht spanish site
var RegexpWgGesuchtSpanish = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/es/`)

//LooksWgGesucht checks if at the very least the url looks like it is from wg-gesucht.de
//i.e. does the url starts with "www.wg-gesucht.de" (+- http(s)://)
func LooksWgGesucht(url string) bool {
	return RegexpWgGesuchtBase.MatchString(url)
}

//IsList checks if the url is a list of offers or requests from wg-gesucht.de
func IsList(url string) bool {
	return RegexpWgGesuchtList.MatchString(url)
}

//IsOfferList checks if the url is a list of offers on wg-gesucht.de
//TODO This one is not possible with regexp package. Need to DIY a bit
func IsOfferList(url string) bool {
	return false
}

//IsRequestList checks if the url is a list of requests on wg-gesucht.de
func IsRequestList(url string) bool {
	return RegexpWgGesuchtRequestList.MatchString(url)
}

//IsEnglish checks if the url is an english version of some page on wg-gesucht.de
func IsEnglish(url string) bool {
	return RegexpWgGesuchtEnglish.MatchString(url)
}

//IsSpanish checks if the url is a spanish version of some page on wg-gesucht.de
func IsSpanish(url string) bool {
	return RegexpWgGesuchtSpanish.MatchString(url)
}

//IsGerman checks if the url is a german version of some page on wg-gesucht.de
func IsGerman(url string) bool {
	return (!IsEnglish(url) && !IsSpanish(url)) //by not being english nor spanish
}
