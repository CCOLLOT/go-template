package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const (
	MESSAGE = "This application is running version 1"
)

var (
	debug          bool
	configFilePath string
)

func InitAndRunCommand() error {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "Run the main process",
	}
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Run the main process",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}
	rootCmd.AddCommand(startCmd)
	return rootCmd.Execute()
}

func Run() error{
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(fmt.Sprintf("%s-%s", t, MESSAGE))
		}
	}
}
