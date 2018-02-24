package classifier

import "testing"

type TestURL struct {
	URL            string
	ExpectedOutput Output
}

var testData = []TestURL{
	TestURL{"https://www.wg-gesucht.de", Output{"wg-gescuht", false, false, "de"}},
	TestURL{"https://www.wg-gesucht.de/en", Output{"wg-gescuht", false, false, "en"}},
	TestURL{"https://www.wg-gesucht.de/en/", Output{"wg-gescuht", false, false, "en"}},
	TestURL{"https://www.wg-gesucht.de/es", Output{"wg-gescuht", false, false, "es"}},
	TestURL{"https://www.wg-gesucht.de/es/", Output{"wg-gescuht", false, false, "es"}},
	TestURL{"http://www.wg-gesucht.de", Output{"wg-gescuht", false, false, "de"}},
	TestURL{"http://www.wg-gesucht.de/en", Output{"wg-gescuht", false, false, "en"}},
	TestURL{"http://www.wg-gesucht.de/es", Output{"wg-gescuht", false, false, "es"}},
	TestURL{"https://www.wg-gesucht.de/somer", Output{"wg-gescuht", false, false, "de"}},
	TestURL{"https://www.wg-gesucht.en", Output{"", false, false, ""}},
	TestURL{"https://www.wg-gesucht.es", Output{"", false, false, ""}},
	TestURL{"", Output{"", false, false, ""}},
	TestURL{"https://", Output{"", false, false, ""}},
	TestURL{"www.wh-gesucht.de", Output{"", false, false, ""}},
	TestURL{"asc;oihwoelkihfoy8239gufobdijpw20huodb1l;kn3w2p", Output{"", false, false, ""}},
	TestURL{"www.wg-gesucht.de", Output{"wg-gescuht", false, false, "de"}},
	TestURL{"www.wg-gesucht.de/en", Output{"wg-gescuht", false, false, "en"}},
	TestURL{"ftp://www.wg-gesucht.de/en/", Output{"", false, false, ""}},
	TestURL{"htttttp://www.wg-gesucht.de/es", Output{"", false, false, ""}},
}

func TestClassify(t *testing.T) {
	for _, aTest := range testData {
		if aTest.ExpectedOutput != Classify(aTest.URL) {
			t.Errorf("It failed at this url: %s", aTest)
		}
	}
}
