package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// uiflag = flag.String("_ui", "gui", "   <选择操作界面> [web] [gui] [cmd]")
	// flag.Parse()
	// addr string
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}
	// addr = fmt.Sprintf("%s:%d", "172.16.9.229", 22)
	client, err := ssh.Dial("tcp", "172.16.9.229:22", &ssh.ClientConfig{
		User:            "xxx",
		Auth:            []ssh.AuthMethod{ssh.Password("***")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	ce(err, "dial")
	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("linux", 32, 160, modes)
	ce(err, "request pty")
	err = session.Shell()
	ce(err, "start shell")
	err = session.Wait()
	ce(err, "return")
	fmt.Println("TESET OK")
	// serialgui.SerialWindow()
}
