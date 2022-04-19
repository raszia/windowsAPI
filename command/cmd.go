package command

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type cmdCommandStruct struct {
	command string
	args    []string
}

func CmdCommand(theCommand string) *cmdCommandStruct {

	return &cmdCommandStruct{
		command: theCommand,
	}

}

func (cmd *cmdCommandStruct) ArgAdd(arg string) *cmdCommandStruct {
	cmd.args = append(cmd.args, arg)
	return cmd
}

func (cmd *cmdCommandStruct) Print() {
	fmt.Println(cmd.command, strings.Join(cmd.args, " "))
}

func (cmd *cmdCommandStruct) Run() error {
	x := exec.Command(cmd.command, cmd.args...)

	var buff, errBuf bytes.Buffer
	x.Stdout = &buff
	x.Stderr = &errBuf
	err := x.Run()
	return err
}
