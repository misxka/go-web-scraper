# Go Web Scraper

A simple CLI-based web scraper written in Go. This application allows you to scrape a webpage, extract all the links, and check their statuses concurrently.

## Features

- Scrapes a webpage and extracts all links.
- Checks the HTTP status of each link.
- Supports concurrency to improve performance.
- Limits the depth of scraping to avoid infinite loops.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/go-web-scraper.git
cd go-web-scraper
```

2. Install dependencies

```sh
go mod tidy
```

3. Build the project

```sh
go build -o bin/
```

## Usage

Run the scraper with the following command:

```sh
./bin/webscraper <URL> <concurrency_limit>
```

  - `<URL>`: The starting URL to scrape.
  - `<concurrency_limit>`: The maximum number of concurrent requests.

### Example

```sh
./bin/webscraper https://example.com 10
```

This will scrape https://example.com with a concurrency limit of 10.

## Configuration
You can modify the following parameters in the code:

- Channel Capacity: The capacity of the tasks channel in scraper.go (default: 1000).
- Scraping Depth: The maximum depth of scraping (default: 1).


## Dependencies
- [Cobra](https://github.com/spf13/cobra) for CLI commands.
- [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html) for parsing HTML.