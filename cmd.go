package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	//cmd := exec.Command("ls", "-lah")

	cmd := exec.Command("/bin/bash", "-c", `ls -lah`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
