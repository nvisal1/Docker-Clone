// +build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> cmd args
// go run maing.go run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Invalid command.")
	}
}

func run() {
	fmt.Printf("running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

