package cmd

import "github.com/spf13/cobra"

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
}

func init() {
	rootCmd.AddCommand(&serverCmd)
	serverCmd.AddCommand(&serverStartCmd)
	serverCmd.AddCommand(&serverStopCmd)
	serverCmd.AddCommand(&clearLogsCmd)
	serverCmd.AddCommand(&reloadQuestsCmd)
}
