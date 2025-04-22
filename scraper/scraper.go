package scraper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Task struct {
	URL   string
	Depth int
}

type ScraperContext struct {
	Tasks   chan Task
	WG      *sync.WaitGroup
	Visited map[string]struct{}
	Mu      *sync.Mutex
}

var i int = 0

func InitScraper(url string, concurrencyLimit int) {
	tasks := make(chan Task, 1000)

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrencyLimit)

	visited := make(map[string]struct{})
	var mu sync.Mutex

	wg.Add(1)
	tasks <- Task{URL: url, Depth: 1}

	for i := 0; i < concurrencyLimit; i++ {
		go func() {
			for task := range tasks {
				mu.Lock()
				if _, ok := visited[task.URL]; ok {
					mu.Unlock()
					wg.Done()
					continue
				}

				visited[task.URL] = struct{}{}
				mu.Unlock()

				sem <- struct{}{}

				fetchUrl(task.URL, task.Depth, &ScraperContext{WG: &wg, Mu: &mu, Tasks: tasks, Visited: visited})

				<-sem
				wg.Done()
			}
		}()
	}

	wg.Wait()
	close(tasks)

	fmt.Println("Scraping complete.")
}

func fetchUrl(url string, depth int, ctx *ScraperContext) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read a response body: ", err)
		return
	}

	if depth <= 1 {
		extractLinks(string(content), depth, ctx)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fmt.Printf("%v: %v\n", url, response.StatusCode)
	} else {
		fmt.Printf("%v: Successful\n", url)
	}
}

func extractLinks(content string, depth int, ctx *ScraperContext) error {
	node, err := html.Parse(strings.NewReader(content))

	if err != nil {
		fmt.Println("Failed to parse HTML content: ", err)
		return err
	}

	var f func(*html.Node, *ScraperContext)
	f = func(node *html.Node, ctx *ScraperContext) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href := attr.Val

					base, err := url.Parse(href)
					if err != nil {
						return
					}

					u, err := url.Parse(href)
					if err != nil {
						continue
					}
					abs := base.ResolveReference(u)

					if abs.Host == base.Host && (abs.Scheme == "http" || abs.Scheme == "https") {
						ctx.Mu.Lock()
						_, found := ctx.Visited[href]
						ctx.Mu.Unlock()

						if !found {
							ctx.WG.Add(1)
							ctx.Tasks <- Task{URL: abs.String(), Depth: depth + 1}
						}
					}
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c, ctx)
		}
	}

	f(node, ctx)
	fmt.Println("Links extracted successfully.")

	return nil
}
