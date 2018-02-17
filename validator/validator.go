package validator

import (
	"regexp"
)

//RegexpWgGesuchtBase is the compiled regex for wg gesucht base
var RegexpWgGesuchtBase = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de`)

//RegexpWgGesuchtList is the compiled regex for wg gesucht list
var RegexpWgGesuchtList = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/(.*)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtRequestList is the compiled regex for wg gesucht request list
var RegexpWgGesuchtRequestList = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/(.*-gesucht)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtEnglish is the compiled regex for wg gesucht english site
var RegexpWgGesuchtEnglish = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/en`)

//RegexpWgGesuchtSpanish is the compiled regex for wg gesucht spanish site
var RegexpWgGesuchtSpanish = regexp.MustCompile(`^(http|https)://www\.wg-gesucht\.de/es`)

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
func IsOfferList(url string) bool {
	//TODO this is not the safest thing ever: if they added a third type of list it will screw me up
	//The reason why I am not using another regexp is because ?! is not supported in regexp package (just google)
	return RegexpWgGesuchtList.MatchString(url) && !RegexpWgGesuchtRequestList.MatchString(url)
}

//TODO IsOffer

//IsRequestList checks if the url is a list of requests on wg-gesucht.de
func IsRequestList(url string) bool {
	return RegexpWgGesuchtRequestList.MatchString(url)
}

//TODO IsRequest

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
	//TODO this is not the safest thing in the world: if they add new language it will screw me up
	//If soneome comes up with a regexp string that does not create false positive nor negative, let me know
	return (LooksWgGesucht(url) && !IsEnglish(url) && !IsSpanish(url)) //by not being english nor spanish
}
