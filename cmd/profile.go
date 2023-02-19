package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/zynko-dev/m2ctl/internal"
)

var profileCmd = cobra.Command{
	Use:   "profile",
	Short: "accesses the profile resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var profileCreateCmd = cobra.Command{
	Use:   "create",
	Short: "creates a new profile",
	Run: func(cmd *cobra.Command, args []string) {
		tag, err := cmd.Flags().GetString("tag")
		if tag == "" || err != nil {
			log.Fatal("The flag tag is mandatory", err)
			return
		}

		host, err := cmd.Flags().GetString("host")
		if host == "" || err != nil {
			log.Fatal("The flag host is mandatory", err)
			return
		}

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		sshKeyFile, _ := cmd.Flags().GetString("ssh-key-file")
		sshKeyPass, _ := cmd.Flags().GetString("ssh-key-pass")

		if (username == "" && password == "") && (sshKeyFile == "" && sshKeyPass == "") {
			log.Fatal("You need to provide either username and password or the path to the ssh-key-fle and the ssh-key-pass")
		}

		internal.CreateRegistryKey()

		if username != "" && password != "" {
			internal.CreateKey(fmt.Sprintf("M2CTL_USERNAME_%s", tag), username)
			internal.CreateKey(fmt.Sprintf("M2CTL_PASSWORD_%s", tag), password)
		}

		if sshKeyFile != "" && sshKeyPass != "" {
			internal.CreateKey(fmt.Sprintf("M2CTL_SSH-KEY-FILE_%s", tag), sshKeyFile)
			internal.CreateKey(fmt.Sprintf("M2CTL_SSH-KEY-PASS_%s", tag), sshKeyPass)
		}

		internal.CreateKey(fmt.Sprintf("M2CTL_HOST_%s", tag), host)

		gitRepoUrl, _ := cmd.Flags().GetString("git-repo-url")
		gitUser, _ := cmd.Flags().GetString("git-username")
		gitEmail, _ := cmd.Flags().GetString("git-email")
		gitSshFile, _ := cmd.Flags().GetString("git-ssh-file")
		gitSshPass, _ := cmd.Flags().GetString("git-ssh-pass")

		if gitRepoUrl != "" && gitUser != "" && gitEmail != "" && gitSshFile != "" && gitSshPass != "" {
			internal.CreateKey(fmt.Sprintf("M2CTL_GIT-REPO-URL_%s", tag), gitRepoUrl)
			internal.CreateKey(fmt.Sprintf("M2CTL_GIT-USER_%s", tag), gitUser)
			internal.CreateKey(fmt.Sprintf("M2CTL_GIT-EMAIL_%s", tag), gitEmail)
			internal.CreateKey(fmt.Sprintf("M2CTL_GIT-SSH-KEY-FILE_%s", tag), gitSshFile)
			internal.CreateKey(fmt.Sprintf("M2CTL_GIT-SSH-KEY-PASS_%s", tag), gitSshPass)
		}

	},
}

var profileDeleteCmd = cobra.Command{
	Use:   "delete",
	Short: "deletes a existing profile identified by the tag",
	Run: func(cmd *cobra.Command, args []string) {
		tag, _ := cmd.Flags().GetString("tag")
		if tag == "" {
			log.Fatal("The flag tag is mandatory")
		}

		internal.DeleteKey(tag)
	},
}

func init() {
	rootCmd.AddCommand(&profileCmd)
	profileCmd.AddCommand(&profileCreateCmd)
	profileCmd.AddCommand(&profileDeleteCmd)

	profileCreateCmd.PersistentFlags().String("host", "", "Host adress for example 127.0.0.1")
	profileCreateCmd.PersistentFlags().String("username", "", "Username to authenticate over ssh for example root")
	profileCreateCmd.PersistentFlags().String("password", "", "Authenticate using the password for given user")
	profileCreateCmd.PersistentFlags().String("ssh-key-file", "", "Path of your loca ssh key file")
	profileCreateCmd.PersistentFlags().String("ssh-key-pass", "", "Password for the ssh key file")
	profileCreateCmd.PersistentFlags().String("git-repo-url", "", "Git repo url to use git as the starting point for the files")
	profileCreateCmd.PersistentFlags().String("git-username", "", "Username of your git user")
	profileCreateCmd.PersistentFlags().String("git-email", "", "Email of your git user")
	profileCreateCmd.PersistentFlags().String("git-ssh-file", "", "Path of your local ssh key file for git online repo")
	profileCreateCmd.PersistentFlags().String("git-ssh-pass", "", "Password of ssh file if one is set")
	profileCreateCmd.PersistentFlags().String("tag", "", "A tag which you can use to refrence the profile in other resources")

	profileDeleteCmd.PersistentFlags().String("tag", "", "Profile with the follwoing tag will be delted")
}
