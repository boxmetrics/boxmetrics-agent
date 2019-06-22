package cmd

import (
	"fmt"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/cli"
	"github.com/spf13/cobra"
	"log"
)

// statusCmd represents the serve command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show boxagent status",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get status")
		status, err := cli.CheckStatus()
		if err != nil {
			log.Fatalln("failed to get status: ", err)
		}

		if status {
			fmt.Println("server is running")
		} else {
			fmt.Println("server is down")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
