package util

import (
	"fmt"
	"os"
	"os/exec"
)

func Runwithcolor(args []string)  {
	_, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{w}

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
