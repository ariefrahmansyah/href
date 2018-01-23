package href

import (
	"context"
	"net/url"
	"strings"
)

func ParseHREF(ctx context.Context, parentURL *url.URL, href string) (*url.URL, error) {
	// if href is starting with '/'
	if href[0] == '/' {
		if len(href) == 1 ||
			len(href) >= 2 && href[1] != '/' {
			newURL := *parentURL
			newURL.Path = href
			return &newURL, nil
		}
	}

	hrefURL, err := url.Parse(href)
	if err != nil {
		return nil, err
	}

	// if href host is empty
	// it may be a relative url
	if hrefURL.Host == "" {
		offset := 0
		if href[0] == '/' {
			offset = 1
			if len(href) > 2 && href[1] == '/' {
				offset = 2
			}
		}
		structure := strings.Split(href[offset:], "/")

		hrefURL.Host = structure[0]
		hrefURL.Path = ""
		if len(structure) > 1 {
			hrefURL.Path = structure[1]
		}
	}

	return hrefURL, nil
}

func IsSameDomain(url1, url2 *url.URL) bool {
	return url1.Hostname() == url2.Hostname()
}
