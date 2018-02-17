# WG-Gesucht.de Scraper

## Goal

- Make a WG-Gesucht scraper library with Golang
- Practice my Golang skills

## What can/will it do?

- Tells you if an URL is an offer, a request, list of offers, or list of requests
	- Input: URL
	- Output: what type of link it is
		- Offer
		- Request
		- Offers list
		- Request list
- Scrap a specific WG-Gesucht.de offer
	- Input: URL of an offer
	- Output: An "Offer" struct that contains all the relevant data, including but not limited to:
		- Which type of offer
			- Flatshare
			- 1 room flat
			- Flats
			- Houses
		- rent
		- what's available in the room/apartment
		- info about the people
	- Error: when the offer is not accessable, or if too many thing broke (i.e. they may have change the site)
- Scrap a specific WG_Gesucht.de request
	- Input: URL of an offer
	- Output: An "Request" struct that contains all the relevant data
- Scrap offers pages into a list of offers
	- Input: URL of that page (list of offers)
	- Output: List of "Offer" struct + node to the next/previous page
- Scrap requests pages into a list of requests
	- Input: URL of that page (list of requests)
	- Output: List of "Request" struct + node to the next/previous page

## What's next

- Follow licenses of dependencies
	- goquery
- Make a crawer-ish thing that uses this to turn HTML into GraphQL API
	- User needs to specify entry point
	- Crawer should crawl through the whole list, page after page
	- User should be able to filter the list
- Check if they have a REST API (i.e. how does the app work)
	- Wireshark
- Check what is the min id of ads
	- right now I have a pretty ugly workaround in IsAd(). It would be useful if I can just ignore any id with less than 2/3/4/5/6 digits
- Write down the research of wg-gesucht urls