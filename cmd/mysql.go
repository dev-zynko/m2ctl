package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zynko-dev/m2ctl/internal"
)

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
	Run: func(cmd *cobra.Command, args []string) {
		device := internal.InitSSHDevice(cmd)

		if device.Debug {
			internal.StdOutPrinter(device.Stdout)
			internal.StdErrPrinter(device.Stderr)
		}
		device.Session.Shell()

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		allowRemoteAcess, _ := cmd.Flags().GetString("allow-remote-login")
		priviliges, _ := cmd.Flags().GetStringArray("privileges")
		tag, _ := cmd.Flags().GetString("tag")

		mysqlPass, _ := internal.GetKey(fmt.Sprintf("M2CTL_MYSQL-PASS_%s", tag))

		device.Stdin.Write([]byte(fmt.Sprintf("mysql -u root -p%s \n", mysqlPass)))
		device.Stdin.Write([]byte(fmt.Sprintf("CREATE USER '%s'@'%s' IDENTIFIED BY '%s'; \n", username, allowRemoteAcess, password)))
		device.Stdin.Write([]byte(fmt.Sprintf("GRANT %s on *.* TO '%s'@'%s' WITH GRANT OPTION; \n", strings.Join(priviliges, ","), username, allowRemoteAcess)))
		device.Stdin.Write([]byte("FLUSH PRIVILEGES; \n"))
		device.Stdin.Write([]byte("exit; \n"))

		for {
			fmt.Println("$")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			fmt.Println(text)

		}
	},
}

var mysqlDeleteUserCmd = cobra.Command{
	Use:   "delete-user",
	Short: "Deletes mysql user. Does not work for root",
}

func init() {
	rootCmd.AddCommand(&mysqlCmd)
	mysqlCmd.AddCommand(&mysqlBackupCmd)
	mysqlCmd.AddCommand(&mysqlUploadBackupCmd)
	mysqlCmd.AddCommand(&mysqlCreateUserCmd)
	mysqlCmd.AddCommand(&mysqlDeleteUserCmd)

	mysqlCreateUserCmd.PersistentFlags().String("username", "", "Username for your new mysql user")
	mysqlCreateUserCmd.PersistentFlags().String("password", "", "Password for your new mysql user")
	mysqlCreateUserCmd.PersistentFlags().String("allow-remote-login", "localhost", "By default remote acces is set to localhost, you can change that by passing your machine IP or % (The % is a wildcard which allows any machine to acces remotly) ")
	mysqlCreateUserCmd.PersistentFlags().StringArray("privileges", []string{"CREATE", "ALTER", "DROP", "INSERT", "UPDATE", "DELETE", "SELECT", "REFERENCES", "RELOAD"}, "Pass in which priviliges your new user should have [CREATE, ALTER, DROP, INSERT, UPDATE, DELETE, SELECT, REFERENCES, RELOAD]")
	mysqlCreateUserCmd.PersistentFlags().Bool("debug", false, "Set debug to true in order to see stdout and stderr on your server")
	mysqlCreateUserCmd.PersistentFlags().String("tag", "", "Refrence to create profile and which data to use")
}
