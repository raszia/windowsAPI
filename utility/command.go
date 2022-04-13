package utility

import (
	"bytes"
	"os/exec"
)

func RunCmd(cmd string, args ...string) error {
	x := exec.Command(cmd, args...)

	var buff, errBuf bytes.Buffer
	x.Stdout = &buff
	x.Stderr = &errBuf
	err := x.Run()
	return err
}
