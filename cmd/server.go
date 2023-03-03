package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zynko-dev/m2ctl/internal"
)

var serverCmd = cobra.Command{
	Use:   "server",
	Short: "Accesses the server resource",
}

var serverStartCmd = cobra.Command{
	Use:   "start",
	Short: "Starts the game server",
}

var serverStopCmd = cobra.Command{
	Use:   "stop",
	Short: "Stops the game server",
}

var clearLogsCmd = cobra.Command{
	Use:   "clear-logs",
	Short: "Clears the server logs",
}

var reloadQuestsCmd = cobra.Command{
	Use:   "reload-quests",
	Short: "Reloads your game quests",
	Run: func(cmd *cobra.Command, args []string) {
		device := internal.InitSSHDevice(cmd)

		if device.Debug {
			internal.StdOutPrinter(device.Stdout)
			internal.StdErrPrinter(device.Stderr)
		}

		files, err := internal.GetKey(fmt.Sprintf("M2CTL_FILES_%s", device.Tag))
		if err != nil {
			log.Fatal("Error geting profile key files", err)
		}

		device.Session.Run(internal.ServerCommands("reload-quests", files))
		device.Session.Close()
	},
}

func init() {
	rootCmd.AddCommand(&serverCmd)
	serverCmd.AddCommand(&serverStartCmd)
	serverCmd.AddCommand(&serverStopCmd)
	serverCmd.AddCommand(&clearLogsCmd)
	serverCmd.AddCommand(&reloadQuestsCmd)

	reloadQuestsCmd.PersistentFlags().Bool("debug", false, "Set debug to true in order to see stdout and stderr on your server")
	reloadQuestsCmd.PersistentFlags().String("tag", "", "Refrence to create profile and which data to use")
}
