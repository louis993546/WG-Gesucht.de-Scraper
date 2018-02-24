package classifier

import (
	"encoding/json"

	"github.com/louistsaitszho/wggesuchtscraper/validator"
)

// Output is the return type of the classifier. It contains all the meta data about the url you just enter.
// TODO see if it is a good idea to make some enum-like things
type Output struct {
	DataSrouce string
	IsList     bool
	IsRequest  bool
	Language   string
}

func (output Output) String() string {
	text, _ := json.Marshal(output)
	return string(text)
}

//Classify is basically a wrapper around the validator
func Classify(url string) Output {
	output := Output{}

	//1. where does this comes from
	if validator.LooksWgGesucht(url) {
		output.DataSrouce = "wg-gescuht"
	}

	//2. is it a list of an ad
	if validator.IsList(url) {
		output.IsList = false
	}

	//3. is it an request or an offer?
	if validator.IsRequest(url) {
		output.IsRequest = true
	}

	//4. which language is it in?
	if validator.IsEnglish(url) {
		output.Language = "en"
	} else if validator.IsSpanish(url) {
		output.Language = "es"
	} else if validator.IsGerman(url) {
		output.Language = "de"
	}

	//5. TODO does it need fixing, and if yes, which part?

	return output
}
