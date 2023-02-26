package internal

import (
	"fmt"
	"io"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type Device struct {
	Stdin   io.Writer
	Stdout  io.Reader
	Stderr  io.Reader
	Session *ssh.Session
	Client  *ssh.Client
	Debug   bool
	Tag     string
}

func InitSSHDevice(cmd *cobra.Command) Device {

	device := Device{}

	device.Debug, _ = cmd.Flags().GetBool("debug")
	device.Tag, _ = cmd.Flags().GetString("tag")
	if device.Tag == "" {
		log.Fatal("Tag is mandatory!")
	}

	host := GetKey(fmt.Sprintf("M2CTL_HOST_%s", device.Tag))
	user := GetKey(fmt.Sprintf("M2CTL_USERNAME_%s", device.Tag))
	pass := GetKey(fmt.Sprintf("M2CTL_PASSWORD_%s", device.Tag))
	keyFile := GetKey(fmt.Sprintf("M2CTL_SSH-KEY-FILE_%s", device.Tag))
	keyFilePass := GetKey(fmt.Sprintf("M2CTL_SSH-KEY-PASS_%s", device.Tag))

	if keyFile != "" {
		device.Session, device.Client = AuthToSSHWithKey(keyFile, keyFilePass)
	} else {
		device.Session, device.Client = AuthToSSHWithCredentials(host, user, pass)
	}

	device.Stdin, device.Stdout, device.Stderr = SshPipe(device.Session)

	return device
}
