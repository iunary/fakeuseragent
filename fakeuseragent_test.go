package fakeuseragent_test

import (
	"testing"

	"github.com/iunary/fakeuseragent"
)

func TestGetUserAgent(t *testing.T) {
	agent := fakeuseragent.GetUserAgent(fakeuseragent.BrowserChrome)
	if agent == "" {
		t.Errorf("Expected non-empty user agent for Chrome, but got an empty string")
	}

	agent = fakeuseragent.GetUserAgent(fakeuseragent.BrowserFirefox)
	if agent == "" {
		t.Errorf("Expected non-empty user agent for Firefox, but got an empty string")
	}

	agent = fakeuseragent.GetUserAgent(fakeuseragent.BrowserMSIE)
	if agent == "" {
		t.Errorf("Expected non-empty user agent for MSIE, but got an empty string")
	}

	agent = fakeuseragent.GetUserAgent(fakeuseragent.BrowserSafari)
	if agent == "" {
		t.Errorf("Expected non-empty user agent for Safari, but got an empty string")
	}

	agent = fakeuseragent.GetUserAgent(fakeuseragent.BrowserOpera)
	if agent == "" {
		t.Errorf("Expected non-empty user agent for Opera, but got an empty string")
	}
}

func TestRandomUserAgent(t *testing.T) {
	agent := fakeuseragent.RandomUserAgent()
	if agent == "" {
		t.Errorf("Expected non-empty random user agent, but got an empty string")
	}

	var browser fakeuseragent.Browser
	for b, userAgent := range fakeuseragent.UserAgents {
		for _, v := range userAgent {
			if agent == v {
				browser = b
				break
			}
		}
	}

	agentFromBrowser := fakeuseragent.GetUserAgent(browser)
	if agentFromBrowser == "" {
		t.Errorf("Expected non-empty user agent for browser %s, but got an empty string", browser)
	}

}
