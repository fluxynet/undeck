package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// Version info
var (
	appCommit = "n/a"
	appBuilt  = "n/a"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "undeck",
		Short: "Undeck card server",
	}

	cmdVersion := &cobra.Command{
		Use:   "version",
		Short: "Get the software version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\nCommit : %v\nBuilt: %v\n", appCommit, appBuilt)
		},
	}
	rootCmd.AddCommand(cmdVersion)

	var server = Server{
		Port: "1337",
	}

	var cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Run:   server.serveCmd,
	}
	rootCmd.AddCommand(cmdServe)

	if err := rootCmd.Execute(); err != nil {
		log.Println("failed to execute command: ", err.Error())
	}
}
