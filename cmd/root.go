package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "m2ctl <resource> <action>",
	Short: "m2ctl by Zynko",
	Long:  "m2ctl is a cli which offers to you an automatic setup of files aswell as build and deploy functionality",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
