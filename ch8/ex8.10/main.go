// Exercise 8.10: HTTP requests may be cancelled by closing the optional Cancel channel in the http.Request struct. Modify the web crawler of Section 8.6 to support cancellation.
// Hint: the http.Get convenience function does not give you an opportunity to customize a Request. Instead, create the request using http.NewRequest, set its Cancel field, then perform the request by calling http.DefaultClient.Do(req).

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

var done = make(chan struct{})

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // aquire token
	list, err := Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

// copied from links
// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// copied from outline2
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	worklist := make(chan []string)  // list of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist
	go func() { worklist <- os.Args[1:] }()

	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Create 20 crawlers goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
