package util

import (
	"os"
	"os/exec"
)

func Runwithcolor(args []string) error {
	_, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{w}

	err = cmd.Run()
	//fmt.Println(err)
	return err
}
