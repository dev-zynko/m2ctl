package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zynko-dev/m2ctl/internal"
)

var deployCmd = cobra.Command{
	Use:   "deploy",
	Short: "Deploys your files to designated server",
	Run: func(cmd *cobra.Command, args []string) {
		device := internal.InitSSHDevice(cmd)

		internal.StdErrPrinter(device.Stderr)
		internal.StdOutPrinter(device.Stdout)
		//device.Session.Shell()
		err := device.Session.Start(fmt.Sprintf(
			"pkg install -y git && "+
				//"git config -–global user.name 'zynko-dev' &&" +
				//"git config --global user.email 'zynko.dev@proton.me'" +
				"%s",
			internal.CompileFliege("4"),
		))

		//_, err := device.Stdin.Write([]byte("pkg install -y git"))
		if err != nil {
			log.Fatal(err)
		}
		//device.Session.Wait()
		fmt.Println("Test")
		for {
			fmt.Println("$")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()
			fmt.Println(text)

		}
	},
}

func init() {
	rootCmd.AddCommand(&deployCmd)
	deployCmd.PersistentFlags().String("tag", "", "Refrence to create profile and which data to use")

}
