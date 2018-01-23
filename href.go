package href

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Link struct {
	URL   *url.URL `json:"-"`
	Text  string   `json:"text,omitempty"`
	HREF  string   `json:"href"`
	Depth int      `json:"-"`
}

func NewLink(ctx context.Context, parentURL *url.URL, text, href string, depth int) Link {
	hrefURL, err := ParseHREF(ctx, parentURL, href)
	if err != nil {
		return Link{}
	}

	link := Link{
		URL:   hrefURL,
		Text:  text,
		HREF:  href,
		Depth: depth,
	}

	return link
}

func (link Link) String() string {
	space := strings.Repeat("\t", link.Depth)
	return fmt.Sprintf("%s(%d) %s - %s", space, link.Depth, link.Text, link.URL)
}

func (link *Link) MarshalJSON() ([]byte, error) {
	url := ""
	if link.URL != nil {
		url = link.URL.String()
	}

	type Alias Link
	return json.Marshal(&struct {
		*Alias
		URLStr string `json:"url,omitempty"`
	}{
		Alias:  (*Alias)(link),
		URLStr: url,
	})
}

func (link Link) IsValidPageLink(ctx context.Context) bool {
	return link.HREF != "" && link.HREF != "/" && link.HREF != "#"
}
