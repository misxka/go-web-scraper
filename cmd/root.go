package cmd

import (
	"github.com/misxka/webscraper/scraper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Web scraper app",
	Short: "A simple Web scraper application",
	Long:  "A simple CLI web scraper application that allows you to check the statuses of all the links of the given page.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		scraper.InitScraper(url)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
