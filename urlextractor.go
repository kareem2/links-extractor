package urlextractor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

var ExtractHrefTagsLinks = false

func ScrapePageUrls(link string) []string {
	var htmlContent = CrawelURL(link)
	var urls = ExtractURLs(htmlContent)
	return urls
}

func ScrapePagesUrls(links []string) []string {

	var urlsMap = make(map[string]int)
	var results = []string{}

	for i := 0; i < len(links); i++ {
		var urls = ScrapePageUrls(links[i])

		for i := 0; i < len(urls); i++ {
			urlsMap[urls[i]] = 1
		}
	}

	for url, _ := range urlsMap {
		results = append(results, url)
	}

	return results
}

func CrawelURL(url string) string {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", UserAgent)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Page crrawling error", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}

func GetUrlsFromFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return make([]string, 0)
	}
	urls := strings.Split(string(data), "\r\n")
	return urls
}

func IsValidURL(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}

	return true
}

func ExtractURLs(html string) []string {
	var normalLinks = []string{}
	var hrefTagsLinks = []string{}

	normalLinks = ExtractLinks(html)
	if ExtractHrefTagsLinks {
		hrefTagsLinks = ExtractHrefTags(html)
	}

	var links = Combine(normalLinks, hrefTagsLinks)
	links = RemoveDuplicates(links)

	return links
}

func ExtractHrefTags(html string) []string {
	var urls = []string{}

	re := regexp.MustCompile(`<a\s+(?:[^>]*?\s+)?href=["'](.*?)"`)
	var hrefTags = re.FindAllStringSubmatch(html, -1)

	for i := 0; i < len(hrefTags); i++ {
		urls = append(urls, hrefTags[i][1])
	}

	return urls
}

func ExtractLinks(html string) []string {
	var urls = []string{}

	re := regexp.MustCompile(`(?i)(?:(?:https?|ftp|file):\/\/|www\.|ftp\.)(?:\([-A-Z0-9+&@#\/%=~_|$?!:,.]*\)|[-A-Z0-9+&@#\/%=~_|$?!:,.])*(?:\([-A-Z0-9+&@#\/%=~_|$?!:,.]*\)|[A-Z0-9+&@#\/%=~_|$])`)

	urls = re.FindAllString(html, -1)

	return urls
}

func RemoveDuplicates(urls []string) []string {
	var urlsMap = make(map[string]int)
	var results = []string{}

	for i := 0; i < len(urls); i++ {
		urlsMap[urls[i]] = 1
	}

	for url, _ := range urlsMap {
		results = append(results, url)
	}

	return results
}

func Combine(arr1 []string, arr2 []string) []string {
	return append(arr1, arr2...)
}
