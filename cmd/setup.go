package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/zynko-dev/m2ctl/internal"
)

var setupCmd = cobra.Command{
	Use:   "setup",
	Short: "Automatic setup of the desired files",
	Run: func(cmd *cobra.Command, args []string) {
		device := internal.InitSSHDevice(cmd)

		if device.Debug {
			internal.StdOutPrinter(device.Stdout)
			internal.StdErrPrinter(device.Stderr)
		}

		sourcePath, _ := cmd.Flags().GetString("source-path")
		if sourcePath == "" {
			log.Fatal("Source path is mandatory")
		}

		var gwg sync.WaitGroup
		/*twg.Add(1)
		fmt.Println("Taring source this can take some time. Pelease wait ....")
		go internal.Tar(sourcePath, &twg)
		twg.Wait()*/
		gwg.Add(1)
		go internal.Gzip("./server-source.tar", "./", &gwg)
		gwg.Wait()

		var wg1, wg2, wg3, wg4, wg5 sync.WaitGroup
		wg1.Add(1)
		go internal.MoveFileOverSFTP("Uploading account.sql", "./public/fliege/account.sql", "/usr/home/account.sql", device.Client, &wg1)
		wg1.Wait()
		wg2.Add(1)
		go internal.MoveFileOverSFTP("Uploading common.sql", "./public/fliege/common.sql", "/usr/home/common.sql", device.Client, &wg2)
		wg2.Wait()
		wg3.Add(1)
		go internal.MoveFileOverSFTP("Uploading log.sql", "./public/fliege/log.sql", "/usr/home/log.sql", device.Client, &wg3)
		wg3.Wait()
		wg4.Add(1)
		go internal.MoveFileOverSFTP("Uploading player.sql", "./public/fliege/player.sql", "/usr/home/player.sql", device.Client, &wg4)
		wg4.Wait()
		wg5.Add(1)
		go internal.MoveFileOverSFTP("Uploading server source", `C:\Users\Zynko\Desktop\m2ctl\server-source.tar.gz`, "/usr/home/server.tar.gz", device.Client, &wg5)
		wg5.Wait()

		//threads, _ := cmd.Flags().GetString("threads")
		mysqlVersion, _ := cmd.Flags().GetString("mysql-version")
		pythonVersion, _ := cmd.Flags().GetString("python-version")

		/*files, err := internal.GetKey(fmt.Sprintf("M2CTL_FILES_%s", device.Tag))
		if err != nil {
			log.Fatal("Error geting profile key files", err)
		}*/

		mysqlPass, err := internal.GetKey(fmt.Sprintf("M2CTL_MYSQL-Pass_%s", device.Tag))
		if err != nil {
			log.Fatal("Error getting profile key mysql pass", err)
		}

		deps := internal.InstallDependencies(pythonVersion, mysqlVersion)
		mysqlConfig := internal.SecureConfigMysql(mysqlPass)
		err = device.Session.Run(fmt.Sprintf("%s%s  \n", deps, mysqlConfig))
		fmt.Print(err)
		device.Session.Close()
	},
}

func init() {
	rootCmd.AddCommand(&setupCmd)
	setupCmd.PersistentFlags().String("threads", "4", "Pass in the amount of threads you want to use for compiling the server source by default it will be 4")
	setupCmd.PersistentFlags().String("source-path", "", "Pass in the path to your source files. If you are using Git it will use the path to navigate though the git repo.")
	setupCmd.PersistentFlags().String("mysql-version", "56", "Adjust the mysql version by default it will use 56. You cann pass 57 or 8 to use newer versions.")
	setupCmd.PersistentFlags().String("python-version", "", "Adjust the python version by default it will be the newest version. You can also pass 27 for python2.7")
	setupCmd.PersistentFlags().Bool("debug", false, "Set debug to true in order to see stdout and stderr on your server")
	setupCmd.PersistentFlags().String("tag", "", "Refrence to create profile and which data to use")
}
