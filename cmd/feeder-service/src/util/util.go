package util

import (
	"fmt"
	"syscall"
)

func KillSystem() {
	err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	if err != nil {
		fmt.Println(err)
	}
}
