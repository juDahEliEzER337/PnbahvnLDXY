// 代码生成时间: 2025-08-30 17:33:15
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "golang.org/x/net/html"
)

// ContentScraper defines the structure for the web content scraper.
type ContentScraper struct {
    URL string
}

// NewContentScraper creates a new instance of ContentScraper with the provided URL.
func NewContentScraper(url string) *ContentScraper {
    return &ContentScraper{URL: url}
}

// ScrapeContent fetches and extracts the text content from the given URL.
func (s *ContentScraper) ScrapeContent() (string, error) {
    // Send a GET request to the URL
    resp, err := http.Get(s.URL)
    if err != nil {
        return "", fmt.Errorf("error fetching URL: %w", err)
    }
    defer resp.Body.Close()

    // Check if the response status code is 200 OK
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    // Parse the HTML content
    node, err := html.Parse(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error parsing HTML: %w", err)
    }

    // Extract text content from the parsed HTML
    var content strings.Builder
    var extractText func(*html.Node)

    extractText = func(n *html.Node) {
        switch n.Type {
        case html.TextNode:
            content.WriteString(n.Data)
        case html.ElementNode:
            if n.Data == "script" || n.Data == "style" {
                for c := n.FirstChild; c != nil; c = c.NextSibling {
                    c.Parent = nil
                }
            } else {
                for c := n.FirstChild; c != nil; c = c.NextSibling {
                    extractText(c)
                }
            }
        }
    }

    extractText(node)

    // Return the extracted content
    return content.String(), nil
}

func main() {
    // Example usage of the ContentScraper
    scraper := NewContentScraper("http://example.com")
    content, err := scraper.ScrapeContent()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Extracted content: ", content)
}
