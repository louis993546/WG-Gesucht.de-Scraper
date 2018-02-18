# Scraper
This is the main part that you will use. Basically it

- Validate if your input url is valid with validator
- Fix url with fixer if url is just a little bit off
- Grab the html file
- Inject all of the results into the to be output struct with injector
- output that as a json string