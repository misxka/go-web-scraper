package cmd

import (
	"fmt"
	"strconv"

	"github.com/misxka/webscraper/scraper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Web scraper app",
	Short: "A simple Web scraper application",
	Long:  "A simple CLI web scraper application that allows you to check the statuses of all the links of the given page.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		concurrencyLimit, err := strconv.Atoi(args[1])

		if err != nil {
			fmt.Printf("Invalid concurrency limit: %v\n", err)
			return
		}

		scraper.InitScraper(url, concurrencyLimit)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
