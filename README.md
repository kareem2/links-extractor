
# URLs Extractor

Crawel pages and extract urls from the HTML content using regular expressions.

**Settings:**
You can make the HTML parser parse the relative links inside href tags by enabling the `ExtractHrefTagsLinks` parameter.

Example:
```
package main

import (
	"crawler/urlextractor"
	"fmt"
)

func main() {

	urlextractor.ExtractHrefTagsLinks = true
	var urls = urlextractor.ScrapePagesUrls([]string{"non-valid url", "https://io.hsoub.com/new"})

	for i := 0; i < len(urls); i++ {
		fmt.Println(urls[i])
	}
}
```

