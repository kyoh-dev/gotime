package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotime",
	Short: "Prints the time for a specified location or timezone",
	Long: `gotime is a simple CLI application that prints the time in a specified timezone or location.
It can also show the difference in time between two timezones.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().Format(time.RFC1123))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
