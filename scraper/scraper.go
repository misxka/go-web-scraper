package scraper

import (
	"fmt"
	"net/http"
)

func InitScraper(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fmt.Printf("%v: %v\n", url, response.StatusCode)
	} else {
		fmt.Printf("%v: Successful\n", url)
	}
}
