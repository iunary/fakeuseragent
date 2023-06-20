package fakeuseragent

import (
	"math/rand"
	"time"
)

type Browser string

const (
	BrowserChrome  Browser = "chrome"
	BrowserEdge    Browser = "edge"
	BrowserMSIE    Browser = "msie"
	BrowserFirefox Browser = "firefox"
	BrowserSafari  Browser = "safari"
	BrowserOpera   Browser = "opera"
)

var (
	UserAgents = map[Browser][]string{
		BrowserChrome: {
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.3",
		},
		BrowserEdge: {
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3 Edge/16.16299",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.3 Edge/16.16275",
		},
		BrowserMSIE: {
			"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
			"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		},
		BrowserFirefox: {
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:54.0) Gecko/20100101 Firefox/54.0",
			"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		},
		BrowserSafari: {
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/601.7.7 (KHTML, like Gecko) Version/9.1.2 Safari/601.7.7",
		},
		BrowserOpera: {
			"Opera/9.80 (Windows NT 6.1; WOW64) Presto/2.12.388 Version/12.18",
			"Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14",
		},
	}

	randSource    = rand.NewSource(time.Now().UnixNano())
	randGenerator = rand.New(randSource)
)

// GetUserAgent retrieves a random user agent for the specified browser.
func GetUserAgent(browser Browser) string {
	agents, ok := UserAgents[browser]
	if !ok {
		return ""
	}

	randomIndex := randGenerator.Intn(len(agents))
	return agents[randomIndex]
}

// RandomUserAgent retrieves a random user agent from the supported browsers.
func RandomUserAgent() string {
	browsers := make([]Browser, 0, len(UserAgents))
	for browser := range UserAgents {
		browsers = append(browsers, browser)
	}

	randomIndex := randGenerator.Intn(len(browsers))
	randomBrowser := browsers[randomIndex]
	return GetUserAgent(randomBrowser)
}
