# href
href will parse href link found in the webpage.

## Usage
Let's say, we found `<a href="/about">About</a>` at http://localhost:8080 webpage.

```go
import "context"
import "net/url"
import "github.com/ariefrahmansyah/href"

ctx := context.Background()
parentURL := url.Parse("http://localhost:8080")
text := "About"
href := "/about"
depth := 0

link := href.NewLink(ctx, parentURL, text, href, depth)

// JSON Result:
// {
//         "text": "About",
//         "href": "/about",
//         "url": "http://localhost:8080/about"
// }
```
