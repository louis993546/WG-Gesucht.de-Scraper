package validator

import (
	"fmt"
	"testing"
)

type TestURL struct {
	URL            string
	LooksWgGesucht bool
	IsList         bool
	IsRequestList  bool
	IsOfferList    bool
	IsEnglish      bool
	IsSpanish      bool
	IsGerman       bool
}

func (tu TestURL) String() string {
	// return "{'min': " + string(ns.Min) + ", 'max': " + string(ns.Max) + "}"
	return fmt.Sprintf("%s, %t, %t, %t, %t, %t, %t, %t", tu.URL, tu.LooksWgGesucht, tu.IsList, tu.IsRequestList, tu.IsOfferList, tu.IsEnglish, tu.IsSpanish, tu.IsGerman)
}

//TODO need more test data to cover as many scenario as possible
//If you want to test more urls, just add another row to here, and specified what condition does it match
var testData = []TestURL{
	TestURL{"https://www.wg-gesucht.de", true, false, false, false, false, false, true},
	TestURL{"https://www.wg-gesucht.de/en", true, false, false, false, true, false, false},
	TestURL{"https://www.wg-gesucht.de/en/", true, false, false, false, true, false, false},
	TestURL{"https://www.wg-gesucht.de/es", true, false, false, false, false, true, false},
	TestURL{"https://www.wg-gesucht.de/es/", true, false, false, false, false, true, false},
	TestURL{"http://www.wg-gesucht.de", true, false, false, false, false, false, true},
	TestURL{"http://www.wg-gesucht.de/en", true, false, false, false, true, false, false},
	TestURL{"http://www.wg-gesucht.de/es", true, false, false, false, false, true, false},
	TestURL{"https://www.wg-gesucht.de/somer", true, false, false, false, false, false, true},
	TestURL{"https://www.wg-gesucht.en", false, false, false, false, false, false, false},
	TestURL{"https://www.wg-gesucht.es", false, false, false, false, false, false, false},
	TestURL{"", false, false, false, false, false, false, false},
	TestURL{"asc;oihwoelkihfoy8239gufobdijpw20huodb1l;kn3w2p", false, false, false, false, false, false, false},
	TestURL{"www.wg-gesucht.de", true, false, false, false, false, false, true},
	TestURL{"www.wg-gesucht.de/en", true, false, false, false, true, false, false},
	TestURL{"www.wg-gesucht.de/en/", true, false, false, false, true, false, false},
	TestURL{"www.wg-gesucht.de/es", true, false, false, false, false, true, false},
	TestURL{"www.wg-gesucht.de/es/", true, false, false, false, false, true, false},
	TestURL{"www.wg-gesucht.de", true, false, false, false, false, false, true},
	TestURL{"www.wg-gesucht.de/en", true, false, false, false, true, false, false},
	TestURL{"www.wg-gesucht.de/es", true, false, false, false, false, true, false},
	TestURL{"www.wg-gesucht.de/somer", true, false, false, false, false, false, true},
	TestURL{"www.wg-gesucht.en", false, false, false, false, false, false, false},
	TestURL{"www.wg-gesucht.es", false, false, false, false, false, false, false},
}

func TestLooksWgGesucht(t *testing.T) {
	for _, url := range testData {
		if LooksWgGesucht(url.URL) != url.LooksWgGesucht {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsList(t *testing.T) {
	for _, url := range testData {
		if IsList(url.URL) != url.IsList {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsRequestList(t *testing.T) {
	for _, url := range testData {
		if IsRequestList(url.URL) != url.IsRequestList {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsOfferList(t *testing.T) {
	for _, url := range testData {
		if IsOfferList(url.URL) != url.IsOfferList {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsEnglish(t *testing.T) {
	for _, url := range testData {
		if IsEnglish(url.URL) != url.IsEnglish {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsSpanish(t *testing.T) {
	for _, url := range testData {
		if IsSpanish(url.URL) != url.IsSpanish {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func TestIsGerman(t *testing.T) {
	for _, url := range testData {
		if IsGerman(url.URL) != url.IsGerman {
			t.Errorf("It failed at this url: %s", url)
		}
	}
}

func BenchmarkLooksWgGesucht(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := LooksWgGesucht(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsList(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsRequestList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsRequestList(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsOfferList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsOfferList(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsEnglish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsEnglish(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsSpanish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsSpanish(testData[0].URL)
		b.Log(answer)
	}
}

func BenchmarkIsGerman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		answer := IsGerman(testData[0].URL)
		b.Log(answer)
	}
}
