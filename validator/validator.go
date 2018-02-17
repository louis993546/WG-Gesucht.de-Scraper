package validator

import (
	"regexp"
)

//RegexpWgGesuchtBase is the compiled regex for wg gesucht base
var RegexpWgGesuchtBase = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)`)

//RegexpWgGesuchtList is the compiled regex for wg gesucht list
var RegexpWgGesuchtList = regexp.MustCompile(`^(|((http|https)://))www\.wg-gesucht\.de/(.*)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtAdOrList is the compiled regex for wg gesucht ad or list (request or offer)
var RegexpWgGesuchtAdOrList = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(.*)(\.)([0-9]{2,7})(\.html)`)

//RegexpWgGesuchtRequestList is the compiled regex for wg gesucht request list
var RegexpWgGesuchtRequestList = regexp.MustCompile(`^(|(http|https)://)www\.wg-gesucht\.de/(.*-gesucht)(\.[0-9]|\.[[0-9][[0-9])(\.\d\.\d\.\d.html)$`)

//RegexpWgGesuchtIsRequest is the compiled regex for a single wg gesucht request
var RegexpWgGesuchtIsRequest = regexp.MustCompile(`(?m)^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(.*)(-gesucht)(\.)([0-9]{1,7})(\.html)`)

//RegexpWgGesuchtAboutFlatshares is the compiled regex for wg gesucht flatshares list/ad
var RegexpWgGesuchtAboutFlatshares = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(/wg-zimmer-in-)(.*)(\.)([0-9]{1,7})(\.html)`)

//RegexpWgGesuchtAbout1RoomFlats is the compiled regex for wg gesucht 1 room flats list/ad
var RegexpWgGesuchtAbout1RoomFlats = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(/1-zimmer-wohnungen-in-)(.*)(\.)([0-9]{1,7})(\.html)`)

//RegexpWgGesuchtAboutFlats is the compiled regex for wg gesucht flats list/ad
var RegexpWgGesuchtAboutFlats = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(/wohnungen-in-)(.*)(\.)([0-9]{1,7})(\.html)`)

//RegexpWgGesuchtAboutHouses is the compiled regex for wg gesucht houses list/ad
var RegexpWgGesuchtAboutHouses = regexp.MustCompile(`^(|(http)(|s)(://))(www.wg-gesucht.de)(/en|/es|)(/haeuser-in-)(.*)(\.)([0-9]{1,7})(\.html)`)

//RegexpWgGesuchtEnglish is the compiled regex for wg gesucht english site
var RegexpWgGesuchtEnglish = regexp.MustCompile(`^(|(http|https)://)www\.wg-gesucht\.de/en`)

//RegexpWgGesuchtSpanish is the compiled regex for wg gesucht spanish site
var RegexpWgGesuchtSpanish = regexp.MustCompile(`^(|(http|https)://)www\.wg-gesucht\.de/es`)

//LooksWgGesucht checks if at the very least the url looks like it is from wg-gesucht.de
//i.e. does the url starts with "www.wg-gesucht.de" (+- http(s)://)
func LooksWgGesucht(url string) bool {
	return RegexpWgGesuchtBase.MatchString(url)
}

//IsList checks if the url is a list of offers or requests from wg-gesucht.de
func IsList(url string) bool {
	return RegexpWgGesuchtList.MatchString(url)
}

//IsAd checks if the url is an ad of offer or request from wg-gesucht.de
func IsAd(url string) bool {
	return RegexpWgGesuchtAdOrList.MatchString(url) && !RegexpWgGesuchtList.MatchString(url)
}

//IsOfferList checks if the url is a list of offers on wg-gesucht.de
func IsOfferList(url string) bool {
	//TODO this is not the safest thing ever: if they added a third type of list it will screw me up
	//The reason why I am not using another regexp is because ?! is not supported in regexp package (just google)
	return RegexpWgGesuchtList.MatchString(url) && !RegexpWgGesuchtRequestList.MatchString(url)
}

//IsOffer checks if the url is a single offer on wg-gesucht.de
func IsOffer(url string) bool {
	return IsAd(url) && !IsRequest(url)
}

//IsRequestList checks if the url is a list of requests on wg-gesucht.de
func IsRequestList(url string) bool {
	return RegexpWgGesuchtRequestList.MatchString(url)
}

//IsRequest checks if the url is a single request on wg-gesucht.de
func IsRequest(url string) bool {
	return RegexpWgGesuchtIsRequest.MatchString(url)
}

//AboutFlatshares checks if the url is about flatshares on wg-gesucht.de (list or ad, request or offer)
func AboutFlatshares(url string) bool {
	return RegexpWgGesuchtAboutFlatshares.MatchString(url)
}

//About1RoomFlats checks if the url is about 1 room flats on wg-gesucht.de (list or ad, request or offer)
func About1RoomFlats(url string) bool {
	return RegexpWgGesuchtAbout1RoomFlats.MatchString(url)
}

//AboutFlats checks if the url is about flats on wg-gesucht.de (list or ad, request or offer)
func AboutFlats(url string) bool {
	return RegexpWgGesuchtAboutFlats.MatchString(url)
}

//AboutHouses checks if the url is about houses on wg-gesucht.de (list or ad, request or offer)
func AboutHouses(url string) bool {
	return RegexpWgGesuchtAboutHouses.MatchString(url)
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
	//TODO this is not the safest thing in the world: if they add new language it will screw me up
	//If soneome comes up with a regexp string that does not create false positive nor negative, let me know
	return (LooksWgGesucht(url) && !IsEnglish(url) && !IsSpanish(url)) //by not being english nor spanish
}
