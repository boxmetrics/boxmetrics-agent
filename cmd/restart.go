package cmd

import (
	"fmt"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/cli"
	"github.com/spf13/cobra"
	"log"
)

// restartCmd represents the serve command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart boxagent",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restarting server...")
		ok, err := cli.StopServer()
		if err != nil {
			fmt.Println("failed to stop server ", err)
		}

		if ok {
			fmt.Println("server stoped")
		}

		pid, err := cli.StartServer()
		if err != nil {
			log.Fatalln("failed to start server ", err)
		}

		fmt.Println("server started with PID ", pid)
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
