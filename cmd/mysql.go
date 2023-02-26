package cmd

import "github.com/spf13/cobra"

var mysqlCmd = cobra.Command{
	Use:   "mysql",
	Short: "Accesses mysql ressource",
}

var mysqlBackupCmd = cobra.Command{
	Use:   "backup-dump",
	Short: "Creates mysql dumps of your current database",
}

var mysqlUploadBackupCmd = cobra.Command{
	Use:   "backup-upload",
	Short: "Uploads mysql backup from specified path",
}

var mysqlCreateUserCmd = cobra.Command{
	Use:   "create-user",
	Short: "Creates a new mysql user",
}

var mysqlDeleteUserCmd = cobra.Command{
	Use:   "delete-user",
	Short: "Deletes mysql user. Does not work for root",
}

func init() {
	rootCmd.AddCommand(&mysqlCmd)
	mysqlCmd.AddCommand(&mysqlBackupCmd)
	mysqlCmd.AddCommand(&mysqlUploadBackupCmd)
	mysqlCmd.AddCommand(&mysqlDeleteUserCmd)
}
