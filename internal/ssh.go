package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/pkg/sftp"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/ssh"
)

func AuthToSSHWithCredentials(host string, username string, password string) (*ssh.Session, *ssh.Client) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatal("Failed to dial connection to ssh", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed creating ssh session", err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm-256color"
	}

	session.RequestPty(term, 100, 100, modes)

	return session, client

}

func AuthToSSHWithKey(key string, password string) (*ssh.Session, *ssh.Client) {
	return nil, nil
}

func MoveFileOverSFTP(srcPath string, dstPath string, client *ssh.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	sftp, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed to create sftp session", err)
	}
	defer sftp.Close()

	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Fatal("Failed opening source-file", err)
	}
	defer srcFile.Close()

	dstFile, err := sftp.Create(dstPath)
	if err != nil {
		log.Fatal("Failed to move source file to server dir", err.Error())
	}
	defer dstFile.Close()

	fi, _ := srcFile.Stat()
	bar := progressbar.DefaultBytes(
		fi.Size(),
		"Uploading server source",
	)

	_, err = io.Copy(io.MultiWriter(dstFile, bar), srcFile)
	if err != nil {
		log.Fatal("Failed copy", err)
	}
}

func SshPipe(session *ssh.Session) (io.Writer, io.Reader, io.Reader) {
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Failed STDIN PIPE", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal("Failed STDOUT PIPE", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal("Failed STDERR PIPE", err)
	}

	return stdin, stdout, stderr
}

func StdOutPrinter(stdout io.Reader) {
	go func() {
		scanner := bufio.NewScanner(stdout)
		for {
			if tkn := scanner.Scan(); tkn {
				rcv := scanner.Bytes()
				raw := make([]byte, len(rcv))
				copy(raw, rcv)
				fmt.Println(string(raw))
			} else {
				if scanner.Err() != nil {
					fmt.Println(scanner.Err())
				} else {
					fmt.Println("io.EOF")
				}
				return
			}
		}
	}()
}

func StdErrPrinter(stderr io.Reader) {
	go func() {
		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
}

func GetOutput(stdout io.Reader) (string, error) { return "", nil }
