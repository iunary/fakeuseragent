# Fake User Agent

`fakeuseragent` is a Go package that provides fake user agent strings for popular web browsers. It can be used for testing, web scraping, or any other scenario where you need to simulate different user agents.

## Installation

To use the package, you need to have Go installed and set up on your machine. Then, you can install the package using the `go get` command:

```bash
go get github.com/iunary/fakeuseragent
```

## Usage
The package provides two main functions for retrieving user agents:

- GetUserAgent: Retrieves a random user agent for the specified browser.
- RandomUserAgent: Retrieves a random user agent from the supported browsers.

Here's an example of how to use the package:

```go
package main

import (
	"fmt"

	"github.com/iunary/fakeuseragent"
)

func main() {
	// Get a random user agent for Chrome
	chromeAgent := fakeuseragent.GetUserAgent(fakeuseragent.BrowserChrome)
	fmt.Println("Chrome User Agent:", chromeAgent)

	// Get a random user agent from the supported browsers
	randomAgent := fakeuseragent.RandomUserAgent()
	fmt.Println("Random User Agent:", randomAgent)
}
```

## Supported Browsers
The package currently supports the following browsers:

- Chrome
- Edge
- MSIE (Internet Explorer)
- Firefox
- Safari
- Opera

Each browser has a predefined set of fake user agent strings that you can retrieve using the provided functions.

## Acknowledgements

This package was inspired by the [fake-useragent](https://github.com/fake-useragent/fake-useragent).
