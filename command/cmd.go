package command

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type cmdCommandStruct struct {
	command string
	args    []string
}
type cmdResultStruct struct {
	stdOut bytes.Buffer
	stdErr bytes.Buffer
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

	err := x.Run()
	return err
}

//run and get output
func (cmd *cmdCommandStruct) RunOutput(ctx context.Context) (*cmdResultStruct, error) {
	ec := exec.Command(cmd.command, cmd.args...)

	result := &cmdResultStruct{}
	ec.Stdout = &result.stdOut
	ec.Stderr = &result.stdErr

	err := ec.Run()

	return result, err
}

func (res *cmdResultStruct) GetStdOut() *bytes.Buffer {
	return &res.stdOut
}

func (res *cmdResultStruct) GetStdErr() *bytes.Buffer {
	return &res.stdErr
}
